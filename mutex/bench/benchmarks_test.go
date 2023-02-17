package main

import (
  "testing"
)

func BenchmarkUseMutex(b *testing.B) {
  for i := 0; i < b.N; i++ {
    useMutex()
  }
}

func BenchmarkUseChannel(b *testing.B) {
  for i := 0; i < b.N; i++ {
    useChannel()
  }
}

//func BenchmarkUseAtomic(b *testing.B) {
  //for i := 0; i < b.N; i++ {
    //useAtomic()
 // }//
//}
