package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Welcome to web requests in golang")

	response, err := http.Get(url) //get a response from the website
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()                     //caller's responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body) //read the response.Body to convert the response to databytes
	if err != nil {
		panic(err)
	}
	content := string(databytes) //the databytes are then converted to string type

	fmt.Println(content)

}
