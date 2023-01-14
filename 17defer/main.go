package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in golang")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three") //The defered print statements are executed last in reverse order
	fmt.Println("Numbers")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
