package main

import "testing"

func BenchmarkGetMinTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMinTriangle(10)
	}
}
func BenchmarkGetMinTriangle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMinTriangle2(10)
	}
}
