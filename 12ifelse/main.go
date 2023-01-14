package main

import "fmt"

func main() {
	userAge := 10
	var result string
	if userAge < 18 { //no brackets on the conditions
		result = "This is not for kids"
	} else if userAge == 18 {
		result = "Welcome to the club"
	} else {
		result = "Good to go"
	}

	fmt.Println(result)

	if num := 6; num < 10 { //you can initialize a variable on the go
		fmt.Printf("%v is less than 10", num)
	} else {
		fmt.Printf("%v is greater than 10", num)
	}
}
