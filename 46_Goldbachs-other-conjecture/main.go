// https://projecteuler.net/problem=46
// It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.
// 9 = 7 + 2×1^2
// 15 = 7 + 2×2^2
// 21 = 3 + 2×3^2
// 25 = 7 + 2×3^2
// 27 = 19 + 2×2^2
// 33 = 31 + 2×1^2
// It turns out that the conjecture was false.
// What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(breakGoldbach())

}

func breakGoldbach() uint64 {

	primes := make(map[uint64]bool)
	for i := uint64(3); ; i++ {
		if isPrime(i) {
			primes[i] = true
		} else {
			if isOdd(i) {
				var flag bool
				for prime := range primes {
					if isInt(math.Sqrt(float64((i - prime) / 2))) {
						flag = true
						break
					}
				}
				if !flag {
					return i
				}
			}
		}
	}
}

func isPrime(number uint64) bool {
	switch {
	case number%2 == 0:
		return false
	default:
		for i := uint64(3); (i * i) <= number; i += 2 {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
}

func isOdd(number uint64) bool {
	if number%2 == 0 {
		return false
	}
	return true
}

func isInt(val float64) bool {
	return val == float64(int(val))
}
