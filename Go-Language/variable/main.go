package main

import "fmt"

func main() {
	name := "keshav"
	fmt.Println("my name is", name)

	var age int = 21
	var height float64 = 60.5
	var decided bool = true

	fmt.Println("my age is", age, "and my height is", height, "and i have decided", decided)

	const pi float64 = 3.14159265359
	fmt.Println("value of pi is", pi)
}
