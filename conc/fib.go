package main

import (
  "fmt"
  "sync"
)

func fibonacci(n int){
  fib := make([]int, n)
  fib[0] = 0
  fib[1] = 1
  for i := 2; i < n; i++ {
    fib[i] = fib[i-1] + fib[i-2]
  }

  fmt.Println(fib)
}

func main() {
  var wg sync.WaitGroup
  wg.Add(2)
  go func() {
    fibonacci(10)
    wg.Done()
  }()
  go func() {
    fibonacci(20)
    wg.Done()
  }()
  wg.Wait()
}
