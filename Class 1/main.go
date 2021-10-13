package main

import "fmt"

// Declaring a variable outside a function and omitting the type since we are initializing it
var position = "Manager"

func main() {
	// Declaring a variable
	var name string
	name = "Fulanito"
	fmt.Println(name)

	// Declaring multiple variables & initializing them
	var (
		lastName, address string = "Pérez", "Calzada X #444"
		age               int    = 30
	)

	fmt.Println(lastName)
	fmt.Println(address)
	fmt.Println(age)

	weight := 88.9
	height := 1.85
	fmt.Println(weight)
	fmt.Println(height)

	fmt.Println(position)
}
