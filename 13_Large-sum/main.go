// https://projecteuler.net/problem=13
// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {

	sum := sumFromFile("number.txt")
	first10 := sum.String()[:10]
	fmt.Println(first10)

}

func sumFromFile(path string) *big.Int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := big.NewInt(0)
	for scanner.Scan() {
		s := scanner.Text()
		n := big.NewInt(0)
		n, ok := n.SetString(s, 10)
		if !ok {
			panic("SetString: error")
		}
		sum.Add(sum, n)
	}
	return sum
}
