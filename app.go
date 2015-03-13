package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

func main () {
  resp, err := http.Get("http://google.com")
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
