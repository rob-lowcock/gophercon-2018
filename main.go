package main

import (
	"fmt"
)

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
	currentX := 0
	currentY := 0

	type coord struct {
		X int
		Y int
	}

	visitedCoords := map[coord]int{coord{0, 0}: 1}
	for _, c := range input {
		switch c {
		case '<':
			currentX--
		case '>':
			currentX++
		case '^':
			currentY++
		case 'v':
			currentY--
		}

		coords := coord{currentX, currentY}
		if _, ok := visitedCoords[coords]; ok {
			visitedCoords[coords]++
		} else {
			visitedCoords[coords] = 1
		}
	}

	return len(visitedCoords)
}
