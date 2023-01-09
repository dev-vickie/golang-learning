package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays!")

	var studentNames [4]string //No initializations (=)

	studentNames[0] = "Victor"
	studentNames[1] = "Bonnie"
	studentNames[3] = "Denno"
	fmt.Println("The student names are: ", studentNames)
	fmt.Println("The students are: ", len(studentNames)) //The length will be 4,regardless of items in the array

	var animalsList = [3]string{"Cat", "Dog", "donkey"} //initialization occurs directly
	fmt.Println("The animals list is :", animalsList, "and length is", len(animalsList))

}
