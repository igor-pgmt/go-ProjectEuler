package main

import (
	"fmt"
	"time"
)

// Get GCD/HCF
func HCF(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	num := 1000000

	start1 := time.Now()
	fmt.Println(Calculate1(num))
	since1 := time.Since(start1)
	fmt.Println(since1)

	start2 := time.Now()
	fmt.Println(Calculate2(num))
	since2 := time.Since(start2)
	fmt.Println(since2)

	start3 := time.Now()
	fmt.Println(Calculate3(num))
	since3 := time.Since(start3)
	fmt.Println(since3)

}

func Calculate1(limit int) (int, int) {
	num3, den7 := 3, 7
	resNum, resDen := 0, 1

	for q := limit; q > 2; q-- {
		p := (num3*q - 1) / den7
		if p*resDen > resNum*q {
			resDen = q
			resNum = p
		}
	}

	return resNum, resDen
}

// This is the fastest function
func Calculate2(limit int) (int, int) {
	num3, den7 := 3, 7
	resNum, resDen := 0, 1
	lowerbound := 2
	den := limit

	for den >= lowerbound {
		num := (num3*den - 1) / den7
		if num*resDen > resNum*den {
			resDen = den
			resNum = num
			lowerbound = resDen / (num3*resDen - den7*resNum)
		}
		den--
	}

	factor := HCF(resNum, resDen)
	resNum /= factor
	resDen /= factor

	return resNum, resDen
}

// Very slow solution, it takes 2,5 min for 1 000 000
func Calculate3(max int) (int, int) {
	var num3, den7 int = 1, max
	for den := 2; den <= max; den++ {
		var num int
		for num = 1; num < den; num++ {
			if 3*den-7*num <= 0 {
				break
			}
		}
		num--
		if num*den7-den*num3 > 0 {
			num3, den7 = num, den
		}
	}
	return num3, den7
}
