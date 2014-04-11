GO OrbitAPI client
==================

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

  info, err := api.Get("info")
  if err != nil {
    log.Fatal("Info error: ", err)
  }

  fmt.Println("Words remaining today: ", info["daily_word_limit"].(float64)-info["words_today"].(float64))

  args := url.Values{}
  args.Add("text", "Jeg liker politikk sa Solberg til VG.")

  info, err = api.Post("tag", args)
  if err != nil {
    log.Fatal("Tag error: ", err)
  }

  fmt.Println(info)
}
```
