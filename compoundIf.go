package main

import "fmt"

func leapYearCalculator(year int) bool {
	var leapYear bool
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				leapYear = true
			} else {
				leapYear = false
			}
		} else {
			leapYear = true
		}
	} else {
		leapYear = false
	}
	return leapYear
}
func main() {
	var year int

	fmt.Println("What year is it")
	fmt.Scanln(&year)
	fmt.Println(year, "being a leap year is", leapYearCalculator(year))
}
