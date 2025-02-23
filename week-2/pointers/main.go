package main

import "fmt"

func main() {
	var a int = 10  //Variable Declarations
	var b *int = &a //Pointer Declarations

	// Print the variables

	println("Address of a is ", &a)
	println("Value of a is ", a)
	println("Value of b is ", b)
	println("Value of *b is ", *b)
	println("Address of b is ", &b)
	println("Address of *b is ", &(*b))
	println("Address of a is ", &a)
	var num int //Variable Declarations integer type
	num = 10
	var ptr *int //Pointer Declarations
	ptr = &num

	// Print the variables

	fmt.Println("Value of num is ", num)
	fmt.Println("Value of ptr is ", ptr)

	// Print the variables using format specifiers

	fmt.Println("data contains through pointer :", *ptr)
}
