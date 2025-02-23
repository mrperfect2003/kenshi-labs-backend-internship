package main //Declaring the package name as main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() { //This is the main function where the execution of the program begins

	fmt.Println("here we will handle the HTTP requests and responses") //Printing the message to the console
	res, err := http.Get("https://fakestoreapi.com/products")          //Sending the GET request to the server
	if err != nil {
		fmt.Println("An error occured", err) //Printing the message to the console
		return
	}
	defer res.Body.Close()                       //Closing the response body
	fmt.Printf("Type of response is: %T/n", res) //Printing the message to the console

	//read the response body
	data, err := ioutil.ReadAll(res.Body) //Reading the response body
	if err != nil {
		fmt.Println("An error occured while reading response", err) //Printing the message to the console
		return
	}
	fmt.Println("Response data is:", string(data)) //Printing the message to the console
}
