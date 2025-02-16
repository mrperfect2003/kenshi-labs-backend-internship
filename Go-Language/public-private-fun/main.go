package main

import "fmt"

func main() {
	name := "keshav"
	age := 21
	height := 60.5

	fmt.Println("Name : ", name, "Age :", age, "Height :", height)

	// Print the type of variables
	fmt.Printf("Name : %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("Type of age: %T\n", age)

	fmt.Printf("Name: %s, Age: %d, Height: %.2f", name, age, height)
}
