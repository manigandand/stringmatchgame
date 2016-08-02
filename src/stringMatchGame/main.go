package main

import (
	"fmt"
	"stringMatchGame/analyse"
)

func main() {

	Insruction := `==============Welcome to StringMatchGame==============
	  It's easy to play, you have chance to conquer our StringMatchGame Robot2.o.
	  Like i said, it's easy beacause you can set target (Maximum no of currect guess either in same position or within the letters).
	  All the best.. Lets play! :)
	`
	fmt.Println(Insruction)
	fmt.Println("Please enter a string only six chars long!")
	fmt.Println()

	winer := analyse.GuessSmarter()

	fmt.Println(winer, " Won")

}
