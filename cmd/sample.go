package main

import (
	"fmt"
	"github.com/dbmedialab/goorbitapi"
	"log"
	"net/url"
)

var (
	apiKey = "Your API key here"
)

func main() {
	api := orbitapi.NewClient(apiKey)

	info := make(chan map[string]interface{}, 1)
	go func() {
		if err := api.AccountInfo(info); err != nil {
			log.Fatal("Info error: ", err)
		}
	}()

	i := <-info
	fmt.Printf("%s, you have %.0f words remaining today.\n", i["name"], i["daily_word_limit"].(float64)-i["words_today"].(float64))

	tag := make(chan *orbitapi.OrbitTag, 1)
	go func() {
		args := &url.Values{}
		args.Add("text", "Jeg liker politikk sa Solberg til Dagbladet.")
		if err := api.ConceptTag(tag, args); err != nil {
			log.Fatal(err)
		}
	}()

	t := <-tag
	fmt.Printf("%#v", t)
}
