package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(fib(int64(1000)))
}

func fib(input int64) int64 {
	max := big.NewInt(input)
	c := big.NewInt(2) // 3rd Fibonacci number
	count := int64(3)  // Fibonacci index counter
	a, b := big.NewInt(1), big.NewInt(1)
	for {
		a.Set(b)
		b.Set(c)
		c.Add(a, b)
		count++
		digitsAmount := CountDigits(c)
		compare := digitsAmount.Cmp(max)
		if compare >= 0 {
			return count
		}
	}
}

func CountDigits(i *big.Int) *big.Int {
	qq := big.NewInt(0).Set(i)
	comparator, count := big.NewInt(0), big.NewInt(0)
	for qq.Cmp(comparator) != 0 {
		qq.Div(qq, big.NewInt(10))
		count.Add(count, big.NewInt(1))
	}
	return count
}
