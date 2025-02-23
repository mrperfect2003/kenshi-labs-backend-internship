package main

func main() {
	var arr [5]int //Declaring an array of size 5
	arr[0] = 1     //Assigning values to the array
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i := 0; i < len(arr); i++ { //Iterating over the array
		println(arr[i]) //Printing the value at index i
	}
}
