package main

import (
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		binarySearch(arr, 5)
	}
}
