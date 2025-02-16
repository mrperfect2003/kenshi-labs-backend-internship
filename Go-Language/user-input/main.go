package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("What is your name?")

	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name)

	//bufio
	fmt.Println("What is going on?")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	fmt.Println("You said:", text)

}
