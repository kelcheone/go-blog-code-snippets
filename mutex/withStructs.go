package main

import (
  "fmt"
  "sync"
)

type Counter struct {
  count int
}

func (c *Counter) Increment() {
  c.count++
}

func (c *Counter) Decrement() {
  c.count--
}

func (c *Counter) Value() int {
  return c.count
}

func main() {
  var wg sync.WaitGroup
  var counter Counter

  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
      counter.Increment()
      wg.Done()
    }()
  }

  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
      counter.Decrement()
      wg.Done()
    }()
  }

  wg.Wait()
  fmt.Println(counter.Value())
}
