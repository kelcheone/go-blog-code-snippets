
package main

import (
  "encoding/json"
  "fmt"
  "os"
)

func main() {
  fmt.Println("Hello, playground")

  data := map[string]interface{}{
    "name": "Kevin",
    "age":  30,
    "address": map[string]interface{}{
      "street": "123 Main St",
      "city":   "New York",
      "state":  "NY",
    },
  }
  fmt.Println("Hello, playground")
  encoder := json.NewEncoder(os.Stdout)

  encoder.SetIndent("", "  ")

  if err := encoder.Encode(data); err != nil {
    panic(err)
  }
  fmt.Println("Hello, playground")
}
