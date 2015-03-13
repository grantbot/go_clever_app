package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

func main () {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "https://api.clever.com/v1.1/sections", nil)
  req.Header.Set("Authorization", "Bearer DEMO_TOKEN")

  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  text, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s", text)
}
