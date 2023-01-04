package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the program"
	fmt.Println(welcome)

	description := "This is a detailed description of the program"
	fmt.Println(description)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of this program:")

}
