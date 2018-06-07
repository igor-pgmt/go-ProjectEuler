package main

import "strconv"

func LargestPalindrome2(n int) (result int) {
	b1, b2 := iPow(10, n-1), iPow(10, n)
	for i := b1; i < b2; i++ {
		for j := i; j < b2; j++ {
			currentProduct := i * j
			if currentProduct > result {
				if isPalindrome2(strconv.Itoa(currentProduct)) {
					result = currentProduct
				}
			}
		}
	}
	return
}

func isPalindrome2(n string) bool {
	return n == reverse2(n)
}

func reverse2(word string) string {
	runes := []rune(word)
	sLen := len(runes)
	for i := 0; i < sLen/2; i++ {
		runes[i], runes[sLen-1-i] = runes[sLen-1-i], runes[i]
	}
	return string(runes)
}
