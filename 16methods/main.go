package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	victor := User{"Victor max", 14, "vic@go.dev", true, true} //create an "object" type User
	fmt.Println(victor)

	fmt.Printf("User details are: %+v\n", victor) //%+v gives details in an ordered manner

	fmt.Printf("Name is: %v ,and email is: %v\n", victor.Name, victor.Email)
	victor.GetStatus() //method is called
	victor.GetNameAge()
}

type User struct { //Use uppercase for the struct's name
	Name       string //Also uppercase for the struct's properties
	Age        int
	Email      string
	Isverified bool
	Islogged   bool
}

func (u User) GetStatus() { //struct is passed to the method
	fmt.Println("Is user logged in:", u.Islogged) //structs properties can then be accessed
}
func (u User) GetNameAge() {
	u.Name = "Victor"
	u.Age = 89 //you can assign new properties on the go
	fmt.Printf("My name is %v and i am %v years old\n", u.Name, u.Age)
}
