package main

import (
  "fmt"
  "sync"
)

type Sub struct {
  vol map[string]int
  mu  sync.Mutex
}

func (s *Sub) getVol() int {
  s.mu.Lock()
  defer s.mu.Unlock()
  return s.vol["vol"]
}

func (s *Sub) setVol(vol int) {
  s.mu.Lock()
  defer s.mu.Unlock()
  s.vol["vol"] = vol
}

func main() {
  s := &Sub{vol: map[string]int{"vol": 0}}
  s.setVol(10)
  fmt.Println(s.getVol())
}
