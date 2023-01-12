package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["TS"] = "Typescript"
	languages["GO"] = "Golang"

	fmt.Println(languages) //print all items in the map

	fmt.Println(languages["TS"]) //print a value given its key

	delete(languages, "JS") //delete an item from the map,given the key
	fmt.Println(languages)

}
