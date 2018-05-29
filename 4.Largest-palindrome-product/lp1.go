package main

func LargestPalindrome1(n int) (result int) {
	b1, b2 := iPow(10, n-1), iPow(10, n)
	for i := b1; i < b2; i++ {
		for j := i; j < b2; j++ {
			if isPalindrome1(i*j) && i*j > result {
				result = i * j
			}
		}
	}
	return
}

func isPalindrome1(n int) bool {
	return n == reverse1(n)
}

func reverse1(n int) (reversed int) {
	for n > 0 {
		reversed = 10*reversed + n%10
		n = n / 10
	}
	return
}
