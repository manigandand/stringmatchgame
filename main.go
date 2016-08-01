package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	// "strconv"
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
	// computer generated random string
	randomStr := RandomString(6)
	turn := "Player"
	totalCorrect, correctPosition := 0, 0
	fmt.Println("characters need to guess, don't reveal this in future:", randomStr)

	fmt.Println("Enter the target point(Eg: 1 to 6): ")
	targetPoint := 0
	fmt.Scan(&targetPoint)
	if targetPoint < 1 || targetPoint > 6 {
		fmt.Println("target point should be greater than 0 and less than or equal to 6")
		// os.Exit(1)
		return
	}

	/*
	   Switch chance to player and computer untill anyone won the match
	*/
	for {
		if turn == "Player" {
			// get user input
			fmt.Println("-------------------------------------------------")
			fmt.Println("Now your turn..!")
			fmt.Print("Enter text to guess: ")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			// fmt.Println(text)
			totalCorrect, correctPosition = stringMatch(randomStr, text)

			if totalCorrect >= targetPoint || correctPosition >= targetPoint {
				fmt.Println("Congrats..! You won the game :)")
				fmt.Println("You have guessed the correct letter but not in the correct position are: ", totalCorrect)
				fmt.Println("You have guessed the correct letter in the correct position are: ", correctPosition)
				break
			} else {
				fmt.Println("Oops, you not guessed the correct characters this time")
				fmt.Println("You have guessed the correct letter but not in the correct position are: ", totalCorrect)
				fmt.Println("You have guessed the correct letter in the correct position are: ", correctPosition)
				turn = "Computer"
			}

		} else if turn == "Computer" {
			// comuter turn
			// let the computer to guess the characters
			fmt.Println("-------------------------------------------------") //
			fmt.Println("Now your computer turn..!")
			computerGuess := RandomString(6)
			fmt.Println("Computer Guess: ", computerGuess)
			totalCorrect, correctPosition = stringMatch(randomStr, computerGuess)
			if totalCorrect >= targetPoint || correctPosition >= targetPoint {
				fmt.Println("Oops sorry, Computer won the game :(")
				fmt.Println("Good luck for next time. Hope you enjoyed this.")
				fmt.Println("Computer guessed the correct letter but not in the correct position are: ", totalCorrect)
				fmt.Println("Computer guessed the correct letter in the correct position are: ", correctPosition)
				break
			} else {
				fmt.Println("Oops, even computer not guessed the correct characters this time")
				fmt.Println("Computer guessed the correct letter but not in the correct position are: ", totalCorrect)
				fmt.Println("Computer guessed the correct letter in the correct position are: ", correctPosition)
				turn = "Player"
			}
		}
	}

}

func stringMatch(randomStr string, text string) (totalCorrect int, correctPosition int) {
	byterandomStr := []byte(randomStr)
	userInput := []byte(text)

	totalCorrect = 0
	correctPosition = 0
	/*
	   Check for the given string matches exactly in the same position
	*/
	lastCorrectChar := make([]string, 6)
	for i := 0; i < len(byterandomStr); i++ {
		// fmt.Println(userInput[i]," ",byterandomStr[i])
		if byterandomStr[i] == userInput[i] {
			correctPosition++
			/*
			   Calculate the strings guesed correctly but not in the same position
			*/
			for j := i + 1; j < len(byterandomStr); j++ {
				// check wheather the current index already searched or not
				if checkIndex(lastCorrectChar, byterandomStr[i]) {
					if byterandomStr[i] == userInput[j] {
						// fmt.Println("r: ",i," ",byterandomStr[i]," u: ",j," ",userInput[j])
						if byterandomStr[j] != userInput[j] {
							totalCorrect++
						}
					}
				}
			}
			lastCorrectChar[i] = string(byterandomStr[i])

		} else {
			/*
			   Calculate the strings guesed correctly but not in the same position
			*/
			for j := 0; j < len(byterandomStr); j++ {
				// check wheather the current index already searched or not
				if checkIndex(lastCorrectChar, byterandomStr[i]) {
					if byterandomStr[i] == userInput[j] {
						// fmt.Println("r: ",i," ",byterandomStr[i]," u: ",j," ",userInput[j])
						if byterandomStr[j] != userInput[j] {
							totalCorrect++
						}
					}
				}
			}
			lastCorrectChar[i] = string(byterandomStr[i])
		}
	}
	return totalCorrect, correctPosition
}
func checkIndex(lastCorrectChar []string, byterandomStr byte) bool {
	for _, element := range lastCorrectChar {
		if element == string(byterandomStr) {
			return false
		}
	}
	return true
}
func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
