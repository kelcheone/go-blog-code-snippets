package main

import (
  "bytes"
  "encoding/json"
  "fmt"
)

func main() {
  data := map[string]interface{}{
    "name": "Kevin",
    "age":  30,
    "address": map[string]interface{}{
      "street": "123 Main St",
      "city":   "New York",
      "state":  "NY",
    },
  }

//  Marshal the data into JSON
  jsonData, err := json.Marshal(data)
  if err != nil {
    panic(err)
  }

//  Create a buffer to store the indented JSON data
  var buf bytes.Buffer

//  Indent the JSON data
  err = json.Indent(&buf, jsonData, "", "  ")
  if err != nil {
    panic(err)
  }

//  Print the indented JSON data

  fmt.Println(buf.String())
}
