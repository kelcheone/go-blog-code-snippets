package main

import "fmt"

func fib(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  close(c)
}

func main() {
 c := make(chan int, 10)
  d := make(chan int, 20)
  go fib(cap(c), c)
  go fib(cap(d), d)
  for i := range c {
    fmt.Println(i)
  }
  for i := range d {
    fmt.Println(i)
  }
}
