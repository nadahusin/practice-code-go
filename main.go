package main

import "fmt"

func main() {
	printSegitiga()
}

func printSegitiga() {

	var i, j, k, rows int

	fmt.Print("Input Number = ")
	fmt.Scanln(&rows)

	for i = rows; i >= 1; i-- {
		for k = 1; k <= (i*2)-1; k++ {
			fmt.Print("*")
		}
		println()

		for j = rows; j >= i; j-- {
			fmt.Print(" ")
		}
	}
}
