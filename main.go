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
	pos := [2]int{0,0}
	houses := [][2]int{pos}

	for _, v := range input {
		switch val := string(v); val {
		case "^":
			pos[1]++
		case ">":
			pos[0]++
		case "<":
			pos[0]--
		case "v":
			pos[1]--
		}

		if !contains(houses, pos) {
			houses = append(houses, pos)
		}
	}

	return len(houses)
}

func contains(haystack [][2]int, needle [2]int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}
