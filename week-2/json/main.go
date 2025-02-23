package main

import (
	"encoding/json"
	"fmt"
)

// defines the structure of the person

type Person struct {
	Name    string `json:"name"`     //`json:"name"` is a struct tag
	Age     int    `json:"age"`      //`json:"age"` is a struct tag
	IsAdult bool   `json:"is_adult"` //`json:"is_adult"` is a struct tag
}

func main() {
	fmt.Println("we are learning JSON")

	//Create a person object

	person := Person{
		Name:    "John Doe",
		Age:     25,
		IsAdult: true,
	}
	fmt.Println("person data is:", person)

	//Convert person into JSON Encoding (Marshalling)

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error in marshalling the data", err)
		return
	}
	fmt.Println("JSON Data is:", string(jsonData))

	//Convert JSON Decoding into person (Unmarshalling)

	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("error in unmarshalling the data", err)
		return
	}
	fmt.Println("Person Data is:", personData)
}
