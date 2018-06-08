// https://projecteuler.net/problem=14
// Longest Collatz sequence
// Problem 14
// The following iterative sequence is defined for the set of positive integers:
// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
// Using the rule above and starting with 13, we generate the following sequence:
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
// Which starting number, under one million, produces the longest chain?
// NOTE: Once the chain starts the terms are allowed to go above one million.

package main

import "fmt"

func main() {
	i := 1000000
	fmt.Println(generateChain(i))
	start, length := findLargestChain(i)
	fmt.Printf("Start number: %d, chain length: %d numbers\n", start, length)
}

func nextCollatz(x int) int {
	switch x % 2 {
	case 0:
		return x / 2
	default:
		return 3*x + 1
	}
}

func generateChain(start int) []int {
	result := []int{start}
	for start != 1 {
		start = nextCollatz(start)
		result = append(result, start)
	}
	return result
}

func findLargestChain(max int) (num, length int) {
	for i := max; i > 1; i-- {
		chain := generateChain(i)
		if len(chain) > length {
			length = len(chain)
			num = i
		}
	}
	return
}
