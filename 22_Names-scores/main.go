package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"unicode"
)

func main() {
	url := "https://projecteuler.net/project/resources/p022_names.txt1"
	content := getFilecontent(url)
	slc := toSlc(content)
	total := allScores(slc)
	fmt.Println(total)
}

// Get file content from the web
func getFilecontent(url string) []uint8 {
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return content
}

// Convert getted content to the sorted slice
func toSlc(content []uint8) []string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	slc := strings.FieldsFunc(string(content), f)
	sort.Strings(slc)
	return slc
}

// Get worth of the word
func wordWorth(word string) (worth int) {
	for _, letter := range word {
		worth += (int(letter) - 64)
	}
	return
}

// Calculate the total of all the name scores
func allScores(names []string) (score int) {
	for k, name := range names {
		score += (k + 1) * wordWorth(name)
	}
	return
}
