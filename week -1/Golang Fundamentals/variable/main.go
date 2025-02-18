package main

import "fmt"

func main() { //Main function
	name := "keshav"                //Variable Declarations
	fmt.Println("my name is", name) //Printing the message to the console

	var age int = 21          //Variable Declarations integer type
	var height float64 = 60.5 //Variable Declarations float type
	var decided bool = true   //Variable Declarations boolean type

	// Print the variables

	fmt.Println("my age is", age, "and my height is", height, "and i have decided", decided)

	// const function used to declare a constant

	const pi float64 = 3.14159265359
	fmt.Println("value of pi is", pi)

}
