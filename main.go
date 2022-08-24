package main

import "fmt"

func main() {
	for i := 6; i >= 1; i-- {
		for k := 1; k <= (i*2)-1; k++ {
			fmt.Print("*")
		}
		println()

		for j := 6; j >= i; j-- {
			fmt.Print(" ")
		}
	}
}
