// https://projecteuler.net/problem=19
// Counting Sundays
// Problem 19

// You are given the following information, but you may prefer to do some research for yourself.

//     1 Jan 1900 was a Monday.
//     Thirty days has September,
//     April, June and November.
//     All the rest have thirty-one,
//     Saving February alone,
//     Which has twenty-eight, rain or shine.
//     And on leap years, twenty-nine.
//     A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

package main

import "fmt"

var months = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func main() {
	fmt.Println(amountOfSundays(1999))
}

func amountOfSundays(year int) (sum int) {
	allCounter := 1                 // Just a day counter to determine day of the week
	for i := 1901; i <= year; i++ { // Iterate years
		if (i%400 == 0) || (i%100 != 0 && i%4 == 0) { // Check if a year is a leap year
			months[1] = 29
		} else {
			months[1] = 28
		}
		for _, days := range months { // Iterate months
			allCounter += days     // Get the first day of the month
			if allCounter%7 == 0 { // If the weekday is sunday
				sum++ // Add 1 to sunday counter
			}
		}
	}
	return
}
