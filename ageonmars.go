package main

import "fmt"

func mars_age(age int) int {
	return age / 687
}

func main() {
	fmt.Println("How old are you: ")
	var age int
	fmt.Scanln(&age)

	age *= 365

	mars := mars_age(age)
	fmt.Println(mars)
}
