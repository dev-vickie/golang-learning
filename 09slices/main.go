package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitlist = []string{"Apple", "Banana", "melon"}
	fmt.Printf("Type is :,%T", fruitlist) //checkin the type = []string

	fruitlist = append(fruitlist, "Mango", "pineapple") //this adds the items to the mentioned slice
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[0:4])
	fmt.Println(fruitlist) //last value is non-inclusive

	//utilize the make keyword to populate a slice
	highscores := make([]int, 4) //make takes the type([]int-slice of int values) and number of items(4) in it
	highscores[0] = 243
	highscores[1] = 150
	highscores[2] = 500
	highscores[3] = 75
	// highscores[4] = 200 //Adding the fifth item brings an error

	highscores = append(highscores, 344, 23, 76, 112, 100) //the append() method realocates the memory based on the items being added

	fmt.Println(highscores)
	//sort package provides methods for sorting out slices
	fmt.Println(sort.IntsAreSorted(highscores)) //returns false

	sort.Ints(highscores)   //sorting the slice
	fmt.Println(highscores) // returns true after sorting

	//removing values from a slice based on index
	var progLanguages = []string{"python", "Dart", "Rust", "C++", "Php", "typescript"}
	fmt.Println(progLanguages)
	index := 0                                                                //index for which the element will be removed
	progLanguages = append(progLanguages[:index], progLanguages[index+1:]...) //remove the item at that index
	fmt.Println(progLanguages)

}
