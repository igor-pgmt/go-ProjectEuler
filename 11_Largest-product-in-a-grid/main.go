// https://projecteuler.net/problem=11
// Largest product in a grid
// Problem 11
// In the 20×20 grid below, four numbers along a diagonal line have been marked in red.
// The product of these numbers is 26 × 63 × 78 × 14 = 1788696.
// What is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20×20 grid?

package main

import "fmt"

func main() {
	fmt.Println(getSum())
}

func getSum() (mult int) {

	// Product of Y-axis (up and down)
	for i := 0; i < (17 * 20); i++ {
		current := numbers[i] * numbers[20+i] * numbers[40+i] * numbers[60+i]
		if current > mult {
			mult = current
		}
	}

	// Product of diagonal and product of X-axis (right and left)
	for i := 0; i < (17*20 - 3); i++ {
		if i%17 == 0 {
			i += 3
			continue
		}
		current := numbers[i] * numbers[21+i] * numbers[42+i] * numbers[63+i]
		if current > mult {
			mult = current
		}
		current = numbers[i] * numbers[1+i] * numbers[2+i] * numbers[3+i]
		if current > mult {
			mult = current
		}
	}

	// Product of opposite diagonal
	for i := 3; i < (17 * 20); i++ {
		if i%20 == 0 {
			i += 3
			continue
		}
		current := numbers[i] * numbers[19+i] * numbers[38+i] * numbers[57+i]
		if current > mult {
			mult = current
		}
	}
	return
}
