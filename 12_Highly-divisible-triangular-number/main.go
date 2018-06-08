package main

import "fmt"

var primes []int

func main() {
	i := 500
	res := GetMinTriangle(i)
	fmt.Println("Triangle number:", res)
}

func EvDiv(k int) int {
	sum := make(map[int]int)
	for i := 0; k > 1; i++ {
		if k%primes[i] == 0 {
			sum[primes[i]]++
			k = k / primes[i]
			i--
		}
	}
	ans := 1
	for _, v := range sum {
		ans *= v + 1
	}
	return ans
}

func getTriangleNums(lines int) (sum int) {
	for i := 1; i <= lines; i++ {
		sum += i
	}
	return
}

func GetMinTriangle(nums int) int {
	tria := getTriangleNums(nums)
	primes = getPrimeNumbers(tria)
	res, i, triangleNums := 0, 0, 0
	for res <= nums {
		i++
		triangleNums = getTriangleNums(i)
		res = EvDiv(triangleNums)
	}

	return triangleNums
}
