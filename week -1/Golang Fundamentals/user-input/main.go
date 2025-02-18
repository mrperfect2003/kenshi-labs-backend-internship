package main

import (
	"bufio"
	"fmt"
	"os"
	"user-input/scan"
)

func main() {
	// fmt.Println("What is your name?")

	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name)

	//bufio
	fmt.Println("What is going on?")    //Printing the message to the console
	reader := bufio.NewReader(os.Stdin) //Creating a reader object
	text, _ := reader.ReadString('\n')  //Reading the input from the user

	fmt.Println("You said:", text) //Printing the message to the console
	scan.PrintMessage("Hello, World!")
}
