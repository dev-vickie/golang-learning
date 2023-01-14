package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in golang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice number is", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("You can open the box")
	case 2:
		fmt.Println("You can move two spots")
	case 3:
		fmt.Println("You can move  three spots")
	case 4:
		fmt.Println("You can move four spots")
	case 5:
		fmt.Println("You can move five spots")
	case 6:
		fmt.Println("You can move six spots,and play again")
	default:
		fmt.Println("What was that?!")
	}
}
