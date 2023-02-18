package main

import (
  "fmt"
  "sync"
)

type Brand struct {
  Mutex sync.Mutex
  Data map[string]string
}

func NewBrand() Brand{
  return Brand{
    Data: make(map[string]string),
  }
}

func (b *Brand) Has(key string) bool {
  b.Mutex.Lock()
  defer b.Mutex.Unlock()
  _, ok := b.Data[key]
  return ok
}

func (b *Brand) Add(key, value string) {
  b.Mutex.Lock()
  defer b.Mutex.Unlock()
  if !b.Has(key) {
    b.Data[key] = value
  }
  return
}

func main() {
  brand := NewBrand()
  brand.Add("name", "Apple")
  fmt.Println(brand.Has("name"))
}
