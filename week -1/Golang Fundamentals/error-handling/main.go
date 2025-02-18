package main

import "fmt"

func divides(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Demoninator cannot be zero")
	}
	return a / b, nil
}
func main() {
	fmt.Println("started Error handling in the function")
	ans, _ := divides(10, 0)
	fmt.Println("Division of two number is ", ans)
}
