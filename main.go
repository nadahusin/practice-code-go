package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println(choosedFilm(7))

}

func choosedFilm(duration int) string {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := ""
	for i := 0; i < 30; i++ {
		a := data[rand.Intn(len(data))]
		b := data[rand.Intn(len(data))]
		if a+b == duration {
			result = strconv.Itoa(a) + " and " + strconv.Itoa(b)
			break
		}
	}

	if result != "" {
		return result
	} else {
		return "Not Found!"
	}
}
