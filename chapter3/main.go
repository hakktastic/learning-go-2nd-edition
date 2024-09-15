package main

import "fmt"

func main() {

	// exercise 1
	fmt.Println("exercise 1")

	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subSlice := greetings[0:2]
	secondSubSlice := greetings[1:4]
	thirdSubSlice := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(subSlice)
	fmt.Println(secondSubSlice)
	fmt.Println(thirdSubSlice)

	// exercise 2
	fmt.Println("exercise 2")

	var message string = "Hi 👩 and 👨"
	fmt.Println(message)
	rs := []rune(message)
	fmt.Println(string(rs[3]))

	// exercise 3
	fmt.Println("exercise 3")

	type Employee struct {
		name     string
		lastName string
		id       int
	}

	firstPerson := Employee{
		"Bob",
		"Marley",
		1,
	}

	secondPerson := Employee{
		name:     "Bob",
		lastName: "Marley",
		id:       2,
	}

	var thirdPerson Employee = Employee{
		name:     "Bob",
		lastName: "Marley",
		id:       3,
	}

	fmt.Println(firstPerson)
	fmt.Println(secondPerson)
	fmt.Println(thirdPerson)
}
