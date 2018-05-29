package main

import (
	"math"
)

func getPrimeNumbers(prop int) (primeNumbers []int) {
	primeNumbers = []int{1, 2, 3, 5, 7} //first prime numbers
	endings := [4]int{1, 3, 7, 9}       //endings for the odd numbers
	currentNum := 0
	for i := 1; len(primeNumbers) < prop; i++ {
		for _, ending := range endings {
			currentNum = ConcatINT(i, ending)
			if shortCheckPrimeNumbers(currentNum, primeNumbers) {
				primeNumbers = append(primeNumbers, currentNum)
				if len(primeNumbers) >= prop {
					break
				}
			}
		}
	}
	return
}

func shortCheckPrimeNumbers(prop int, primes []int) bool {
	sqrt := int(math.Sqrt(float64(prop)))
	for i := 2; primes[i] <= sqrt; i++ {
		if prop%primes[i] == 0 {
			return false
		}
	}
	return true
}

func ConcatINT(a, b int) int {
	lenB := GetLen(b)
	exp := int(math.Pow(10, float64(lenB)))
	return exp*a + b
}

func GetLen(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return
}

func iPow(a, b int) int {
	var result int = 1

	for 0 != b {
		if 0 != (b & 1) {
			result *= a
		}
		b >>= 1
		a *= a
	}

	return result
}
