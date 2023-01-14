package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	victor := User{"Victor max", 14, "vic@go.dev", true, false} //create an "object" type User
	fmt.Println(victor)

	fmt.Printf("User details are: %+v\n", victor) //%+v gives details in an ordered manner

	fmt.Printf("Name is: %v ,and email is: %v", victor.Name, victor.Email)
}

type User struct { //Use uppercase for the struct's name
	Name       string //Also uppercase for the struct's properties
	Age        int
	Email      string
	Isverified bool
	Islogged   bool
}
