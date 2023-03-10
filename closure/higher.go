package main

import "fmt"

func myfunc(f func(int) int) {
    fmt.Println(f(5))
}

func myFunc2() func (int) int{
  var sum int
  return func (x int) int {
    sum += x
    return sum
  }
}

func main() {
    myfunc(func(x int) int {
        return x * x
    })

    f := myFunc2()
    for i := 0; i < 10; i++ {
      fmt.Println(f(i))
    }
}
