// https://projecteuler.net/problem=10
// Summation of primes
// Problem 10
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

package main

import "fmt"

func main() {
	fmt.Println(sum(2000000))
}

func sum(max int) (sum int) {
	primes := getPrimeNumbers(max)
	fmt.Println(primes[0], primes[len(primes)-1])
	for _, v := range primes {
		sum += v
	}
	return
}
