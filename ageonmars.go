package main

import "fmt"

// Dividing earth days by one year on mars to get mars age
func mars_age(age int) int {
	return age / 687
}

func main() {
	fmt.Println("How old are you: ")
	var age int
	fmt.Scanln(&age)

	// Converting one year earth to day on earth
	age *= 365

	mars := mars_age(age)
	fmt.Println(mars)
}
