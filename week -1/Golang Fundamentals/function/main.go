package main

import "fmt" //Importing fmt package because we are using fmt.Println

func add(a int, b int) int { //Function to add two numbers
	return a + b //Returning the sum of two numbers
}

func main() { //Main function
	ans := add(2, 3)                         //Calling the add function
	fmt.Println("add of two numbe is ", ans) //Printing the result
}
