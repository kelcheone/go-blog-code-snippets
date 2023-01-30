package main

import "testing"

func BenchmarkFuncToBenchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcToBenchmark()
	}
}
