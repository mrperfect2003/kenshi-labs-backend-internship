// Stack example
package main

import "fmt"

func main() {
	// Stack example
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println("Stack:", stack)
	stack = stack[:len(stack)-1] // Pop
	fmt.Println("After pop:", stack)
}
