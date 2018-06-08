// https://projecteuler.net/problem=16
// Power digit sum
// Problem 16
// 215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

// What is the sum of the digits of the number 21000?

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	fmt.Println(powerDigitSum(2, 1000))

}

func powerDigitSum(x, y int) (sum int) {
	pow := new(big.Int).Exp(big.NewInt(int64(x)), big.NewInt(int64(y)), nil).String()

	for _, v := range pow {
		ncum, _ := strconv.Atoi(string(v))
		sum += ncum
	}

	return
}
