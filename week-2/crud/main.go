package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning CRUD operations") //Printing the message to the console

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") //Fetching the data from the API
	if err != nil {
		fmt.Println("Error fetching the http get operation", err) //Printing the message to the console
		return
	}
	defer res.Body.Close() //Closing the response body

	//Checking the status code of the response

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error is getting in the response", res.StatusCode) //Printing the message to the console
		return
	}

	data, err := ioutil.ReadAll(res.Body) //Reading the data from the response body
	if err != nil {
		fmt.Println("Error reading the response body:v", err) //Printing the message to the console
		return
	}
	fmt.Println(string(data)) //Printing the data to the console

}
