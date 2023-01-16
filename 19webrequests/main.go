package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Welcome to web requests in golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //caller's responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)

}
