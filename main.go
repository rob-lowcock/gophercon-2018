package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Sainsbury's GopherCon Challenge 2018!")
	fmt.Println("Calling your function...")

	input := "<<^>vv>^<<v"

	houses := CalculateHouses(input)

	if houses == 8 {
		fmt.Println("Good start! 👍  Run the tests with ginkgo to check everything works")
		return
	}

	fmt.Println("Oh dear - looks like something is wrong with your algorithm 🙁")
}

func CalculateHouses(input string) int {
	return 0
}
