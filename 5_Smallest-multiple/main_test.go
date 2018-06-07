package main

import "testing"

func BenchmarkEvDiv1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvDiv1(5)
	}
}
func BenchmarkEvDiv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvDiv2(5)
	}
}
func BenchmarkEvDiv3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvDiv3(5)
	}
}
