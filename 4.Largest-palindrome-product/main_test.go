package main

import "testing"

func BenchmarkLargestPalindrome1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindrome1(3)
	}
}
func BenchmarkLargestPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindrome2(3)
	}
}
