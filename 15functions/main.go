package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greetings("victor")
	results := addAll(2, 4, 5, 3, 33, 6, 1)

	fmt.Println(results)

}
func greetings(name string) {
	fmt.Println("Hello ", name)
}
func addAll(values ...int) int { //accept only int values
	total := 0
	for _, val := range values { //get each val in the values and add to total
		total += val
	}
	return total //return final total - type int
}
