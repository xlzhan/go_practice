package main

import "fmt"

func main() {
	monthes := map[string]int{
		"Jan": 31, "Feb": 29, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
	}
	fmt.Println(monthes)
	year := 0
	for _, days := range monthes {
		year += days
	}
	fmt.Printf("Days of year: %d\n", year)
}
