package main

import "testing"

func BenchmarkFuncToWithIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcToWithIO()
	}
}

func BenchmarkFuncToWithBufio(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcToWithBufio()
	}
}
