// https://projecteuler.net/problem=9
// Special Pythagorean triplet
// Problem 9
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import "fmt"

func main() {
	fmt.Println(searchTriplet(1000))
}

func searchTriplet(num int) (int, int, int, int) {
	for a := 3; a < (num-3)/3; a++ {
		for b := a + 1; b < (num-a)/2; b++ {
			c := num - a - b
			if a*a+b*b == c*c {
				return a, b, c, a * b * c
			}
		}
	}
	return 0, 0, 0, 0
}
