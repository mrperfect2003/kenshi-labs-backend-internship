// file named main.go in the calculator directory
package main

import (
	"fmt" //Importing fmt package because we are using fmt.Println
)

func main() {
	var operator string //Variable Declarations
	var num1, num2 float64

	fmt.Println("Enter first number:") //Printing the message to the console
	fmt.Scanln(&num1)                  //Taking input from the user

	fmt.Println("Enter operator (+, -, *, /):")
	fmt.Scanln(&operator)

	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)

	switch operator { //Switch case to perform the operation based on the operator
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}
}
