package main

import (
  "fmt"
  "sync"
)

type Person struct {
  Name string
  Age  int
}

func (p *Person) ChangeName(name string, mutex *sync.Mutex) {
  mutex.Lock()
  p.Name = name
  mutex.Unlock()
}

func (p *Person) ChangeAge(age int, mutex *sync.Mutex) {
  mutex.Lock()
  p.Age = age
  mutex.Unlock()
}

func main() {
  var mutex sync.Mutex
  person := Person{
    Name: "Kevin",
    Age:  25,
  }

  go person.ChangeName("John", &mutex)
  go person.ChangeAge(30, &mutex)

  fmt.Println(person)
}
