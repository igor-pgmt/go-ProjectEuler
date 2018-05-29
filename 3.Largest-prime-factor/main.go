// https://projecteuler.net/problem=3
// Largest prime factor
// Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

var number = 600851475143

func main() {

	fmt.Println(LargestPrime1(number)) // 6857
	fmt.Println(LargestPrime2(number)) // 6857

}

func LargestPrime1(num int) int {
	r := num
	for i := 2; i < num; i++ {
		if num%i == 0 {
			r = num / i
			break
		}
	}

	if r == num {
		return num
	}

	return LargestPrime1(r)
}

// Or we can create a separate function for the largest Composite number
func lcn(num int) int {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return num / i
		}
	}
	return num
}

func LargestPrime2(num int) int {
	r := lcn(num)
	if r == num {
		return num
	}
	return LargestPrime2(r)
}
