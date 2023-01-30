package main

import "fmt"

func main() {
    x := make(map[string]map[string]int)

    x["foo"] = make(map[string]int)
    x["foo"]["bar"] = 1

    x["bar"] = make(map[string]int)
    x["bar"]["foo"] = 2

    // Iterating over a nested map
    for key, value := range x {
        fmt.Println(key, value)
    }
  for key, value :=range x{
        for key2, value2 := range value {
            fmt.Println(key, key2, value2)
        }
    }
}
