// https://projecteuler.net/problem=4
// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import "fmt"

func main() {
	fmt.Println("LargestPalindrome1 =", LargestPalindrome1(3))
	fmt.Println("LargestPalindrome2 =", LargestPalindrome2(3))
}

func iPow(a, b int) int {
	var result = 1

	for 0 != b {
		if 0 != (b & 1) {
			result *= a
		}
		b >>= 1
		a *= a
	}

	return result
}
