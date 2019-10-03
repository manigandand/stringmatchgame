package main

import (
	"fmt"
	"stringmatchgame/analyse"
	"stringmatchgame/random"
)

func main() {
	// refill the input string to generate random string for computer move
	random.RefillInputFile()

	Insruction := `==============Welcome to StringMatchGame==============
	  It's easy to play, you have chance to conquer our StringMatchGame Robot2.o.
	  Like i said, it's easy beacause you can set target (Maximum no of currect guess either in same position or within the letters).
	  All the best.. Lets play! :)
	`
	fmt.Println(Insruction)
	fmt.Println("Please enter a string only six chars long!")
	fmt.Println()

	winer := analyse.GuessSmarter()

	if winer == "You" {
		fmt.Println("Congrats..! You won the game :)")
	} else if winer == "Computer" {
		fmt.Println("Oops sorry, Computer won the game :(")
		fmt.Println("Good luck for next time. Hope you enjoyed this.")
	}

}
