package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
)

func main () {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "https://api.clever.com/v1.1/sections", nil)
  req.Header.Set("Authorization", "Bearer DEMO_TOKEN")

  type Section struct {
    //Uri string `json:"uri"`
    Data struct {
      Course_Name string `json:"course_name"`
      Students []string `json:"students"`
    } `json:"data"`

  }

  type Section_Data struct {
    Data []Section `json:"data"`
  }


  // Get response
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  // Read response
  body, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    log.Fatal(err)
  }

  // Unmarshal response onto our structs
  var data Section_Data
  err = json.Unmarshal(body, &data)
  if err != nil {
    fmt.Printf("%T\n%s\n%#v\n",err, err, err)
    switch v := err.(type){
    case *json.SyntaxError:
      fmt.Println(string(body[v.Offset-40:v.Offset]))
    }
  }
  fmt.Printf("%s", data)

  //for i, section := range data.Sections{
    //fmt.Printf("%d: %s %s\n", i, section.Students[0])
  //}
  //fmt.Printf("%s", text.data)
}
