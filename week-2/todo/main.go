package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("learning TODO in golang")

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

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding the response body Of TODO decoding", err)
		return
	}

	fmt.Println("todo", todo)

}
