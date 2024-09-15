package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// exercise 1
	intSlice2 := make([]int, 100)

	for i := range intSlice2 {
		intSlice2[i] = rand.Intn(100)
	}

	fmt.Println(intSlice2)

	// exercise 2
	for i, v := range intSlice2 {

		switch {
		case v%2 == 0:
			fmt.Println("index:", i, "value:", v, " --> Two!")
		case v%3 == 0:
			fmt.Println("index:", i, "value:", v, " --> Three!")
		case v%2 == 6:
			fmt.Println("index:", i, "value:", v, " --> Six!")
		default:
			fmt.Println("index:", i, "value:", v, " --> Never mind")
		}
	}

	// exercise 3
	var total int

	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Println(total)
	}

	fmt.Println("final total: ", total)
}
