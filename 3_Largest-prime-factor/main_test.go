package main

import (
	"reflect"
	"runtime"
	"testing"
)

type testCase struct {
	number int
	result int
}

var testCases = []testCase{
	{100, 5},
	{101, 101},
	{13195, 29},
	{600851475143, 6857},
}

func TestLargestPrime(t *testing.T) {
	solutions := []interface{}{LargestPrime1, LargestPrime2}
	for _, solution := range solutions {
		for _, testCase := range testCases {
			result := solution.(func(int) int)(testCase.number)
			if result != testCase.result {
				t.Errorf("%s\n Argument: %d Result: %d Should be: %d\n", runtime.FuncForPC(reflect.ValueOf(solution).Pointer()).Name(), testCase.number, result, testCase.result)
			}
		}
	}
}

func BenchmarkLargestPrime1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPrime1(number)
	}
}

func BenchmarkLargestPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPrime2(number)
	}
}
