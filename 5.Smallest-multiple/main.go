// https://projecteuler.net/problem=5
// Smallest multiple
// Problem 5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	fmt.Println(EvDiv1(20))
	fmt.Println(EvDiv2(20))
	fmt.Println(EvDiv3(20))
}

func EvDiv1(m int) int {
	for i := 1; ; i++ {
		z := 0
		for j := 1; j <= m; j++ {
			z += i % j
		}
		if z == 0 {
			return i
		}
	}
	return m
}

func EvDiv2(m int) int {
	for i := m; ; i += m {
		z := 0
		for j := 2; j <= m; j++ {
			z += i % j
		}
		if z == 0 {
			return i
		}
	}
	return m
}

// This is the fastest function
func EvDiv3(k int) *big.Int {
	N := big.NewInt(1)
	check := true
	limit := math.Sqrt(float64(k))
	p := getPrimeNumbers(k)
	for i := 1; p[i] <= k; i++ {
		a := 1
		if check {
			if p[i] <= int(limit) {
				c := math.Log(float64(k)) / math.Log(float64(p[i]))
				a = int(math.Floor(c))
			} else {
				check = false
			}
		}
		N = big.NewInt(0).Mul(N, big.NewInt(int64(iPow(p[i], a))))
	}
	return N
}
