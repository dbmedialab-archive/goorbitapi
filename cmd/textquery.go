// Copyright 2014 DB Medialab.  All rights reserved.
// License: MIT

// $ ORBIT_API_KEY=YOUR-KEY go run textquery.go -text "Jeg liker politikk sa Solberg til Dagbladet." -sentences 3
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dbmedialab/goorbitapi"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	apiKey    string
	text      string
	sentences int
)

func init() {
	apiKey = os.Getenv("ORBIT_API_KEY")
	flag.StringVar(&text, "text", "", "Text to use in your query")
	flag.IntVar(&sentences, "sentences", 1, "Number of sentences to extract from Wikipedia")
}

func main() {
	flag.Parse()

	if apiKey == "" || text == "" {
		flag.Usage()
		return
	}

	api := orbitapi.NewClient(apiKey)

	tag := make(chan *orbitapi.OrbitTag, 1)
	go func() {
		args := &url.Values{}
		args.Add("text", text)
		if err := api.ConceptTag(tag, args); err != nil {
			log.Fatal(err)
		}
	}()

	t := <-tag
	if len(t.Entities) < 1 {
		fmt.Println("Nothing!")
		return
	}

	fmt.Printf("What we think we know about that:\n")

	desc := make(chan string, len(t.Entities))
	for _, t := range t.Text[0].([]interface{}) {
		if item, ok := t.(map[string]interface{}); ok {
			go queryWiki(desc, item)
		}
	}

	for i := 0; i < len(t.Entities); i++ {
		fmt.Printf(<-desc)
	}
	fmt.Println()
}

func queryWiki(desc chan<- string, item map[string]interface{}) {
	url := fmt.Sprintf("http://no.wikipedia.org/w/api.php?format=json&action=query&titles=%s&prop=extracts&exsentences=%d&explaintext", item["entity"].(string), sentences)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	var r string
	// Data is held in an unknown key
	for _, v := range data["query"].(map[string]interface{})["pages"].(map[string]interface{}) {
		if e, ok := v.(map[string]interface{})["extract"].(string); ok {
			r = e
		} else {
			r = "Nothing yet!"
		}
		break
	}
	desc <- fmt.Sprintf("\n%s\n\t%s\n", item["text"], r)
}
