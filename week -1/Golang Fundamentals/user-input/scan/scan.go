package scan

import "fmt"

func PrintMessage(message string) { //Function to print the message
	fmt.Println(message) //Printing the message to the console

	// Scan function is used to take the input from the user
	fmt.Println("What is your name?")

	var name string //Variable Declarations
	fmt.Scan(&name) //Taking the input from the user
	fmt.Println("Hello, Mr.", name)

}
