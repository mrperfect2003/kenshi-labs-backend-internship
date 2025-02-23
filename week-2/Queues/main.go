package main

import "fmt"

func main() {

	queue := []int{} // Queue example
	queue = append(queue, 10)
	queue = append(queue, 20)
	queue = append(queue, 30)
	fmt.Println("Queue:", queue)
	queue = queue[1:] // Dequeue
	fmt.Println("After dequeue:", queue)
}
