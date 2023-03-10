package main

import "fmt"

func fibonacci() func() int {
  x, y := 0, 1
  return func() int {
    x, y = y, x+y
    return x
  }
}

func main() {
    f := fibonacci()
    for i :=f(); i < 1000; i = f() {
        fmt.Println(i)
    }
}
