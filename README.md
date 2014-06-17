GO OrbitAPI client
==================

[![Build Status](https://travis-ci.org/dbmedialab/goorbitapi.svg)](https://travis-ci.org/dbmedialab/goorbitapi)

Go client for Orbit API - http://orbitapi.com/


```GO
package main

import (
  "fmt"
  "github.com/dbmedialab/goorbitapi"
  "log"
  "net/url"
)

var (
  apiKey = "Your API key"
)

func main() {
  api := orbitapi.NewClient(apiKey)

  go func() {
    if err := api.AccountInfo(); err != nil {
      log.Fatal("Info error: ", err)
    }
  }()

  r := <-api.Result
  result := r.(map[string]interface{})

  fmt.Println("Words remaining today: ", result["daily_word_limit"].(float64)-result["words_today"].(float64))

  go func() {
    args := &url.Values{}
    args.Add("text", "Jeg liker politikk sa Solberg til Dagbladet.")
    if err := api.ConceptTag(args); err != nil {
      log.Fatal(err)
    }
  }()

  r = <-api.Result
  result = r.(*OrbitTag)

  fmt.Printf("%#v", result)
}
```
