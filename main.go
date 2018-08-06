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

// Fastest solution in 28.03s 🎉
func CalculateHouses(input string) int {
	var c[2]int;g:=map[[2]int]int{c:0};for _,f:=range input{
		c[f/64]+=int(f+10)&8/4-1;g[c]++
	}
	return len(g)
}
