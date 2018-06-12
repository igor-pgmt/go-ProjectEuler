// Coded triangle numbers
// https://projecteuler.net/problem=42
// The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
// By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.
// Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?

package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	url := "https://projecteuler.net/project/resources/p042_words.txt"
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	// Reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	check(err)

	// Get the HTML code as a string
	r := csv.NewReader(strings.NewReader(string(html)))

	// Parse clean CSV content
	words, err := r.Read()

	// Answers counter
	counter := 0

	// Iterating words
	for _, word := range words {

		v := strings.ToLower(word) // Converting all words to lowercase
		sum := 0                   // The sum of the values of the letters in the word
		for i := 0; i < len(v); i++ {
			num := int(v[i]) - 96 // Getting letter number from 1 to 26
			sum += num            // The sum of all letters
		}

		d := 1 + 4*2*sum // Discriminant

		// If the discriminant less then zero, there is no answer
		if d < 0 {
			continue
		}

		// Calculating the roots
		sqrt := math.Sqrt(float64(d))
		x1 := (-1 + sqrt) / 2
		x2 := (-1 - sqrt) / 2

		// If x1 or x2 is bigger than 0 and is integer, it's an answer
		if (x1 > 0 && x1 == float64(int64(x1))) || (x2 > 0 && x2 == float64(int64(x2))) {
			var answer float64
			if x1 > x2 {
				answer = x1
			} else {
				answer = x2
			}
			fmt.Printf("Word: %v, Index number: %v\n", v, answer)
			counter++
		}
	}

	fmt.Println("Total suitable words:", counter)
}
