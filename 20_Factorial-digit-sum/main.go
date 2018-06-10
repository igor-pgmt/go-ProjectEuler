package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	i := *big.NewInt(100)
	fmt.Println(sumOfDigits(factorial(i)))
}

func factorial(n big.Int) *big.Int {
	if n.Int64() == big.NewInt(1).Int64() {
		return &n
	}
	sub := big.NewInt(0)
	sub = sub.Sub(&n, big.NewInt(1))
	return n.Mul(&n, factorial(*sub))
}

func sumOfDigits(n *big.Int) (sum uint64) {
	nstr := n.String()
	fmt.Println(nstr)
	for _, digit := range nstr {
		d, _ := strconv.Atoi(string(digit))
		sum += uint64(d)
	}
	return
}
