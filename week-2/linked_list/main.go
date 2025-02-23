// Simple Linked List Example
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	// Creating nodes
	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}

	// Linking nodes
	n1.next = n2
	n2.next = n3

	// Display linked list
	current := n1
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
