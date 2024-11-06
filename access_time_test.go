package main

import (
	"testing"
)

var mylist = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var mymap = map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10}

// Benchmark for array access
func BenchmarkArrayAccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mylist[5] // Access a middle element repeatedly
	}
}

// Benchmark for map access
func BenchmarkMapAccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mymap[5] // Access a middle key repeatedly
	}
}
