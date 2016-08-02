package analyse

import (
	"bufio"
	"fmt"
	"os"
	"stringMatchGame/engine"
	"stringMatchGame/random"
)

func GuessSmarter() string {
	// computer generated random string
	randomStringLength := 6
	randomStr := random.RandomString(randomStringLength)

	winer := ""
	turn := "Player"
	totalCorrect, correctPosition := 0, 0
	fmt.Println("characters need to guess, don't reveal this in future:", randomStr)

	targetPoint := 0

	for {
		if targetPoint < 1 || targetPoint > 6 {
			fmt.Println("Enter the target point(Eg: 1 to 6): ")
			fmt.Scan(&targetPoint)
			if targetPoint < 1 || targetPoint > 6 {
				fmt.Println("target point should be greater than 0 and less than or equal to 6")
			}
		} else {
			break
		}
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
			text := ""
			// get the correct length of string
			for {
				text, _ = reader.ReadString('\n')
				if len(text) != 7 {
					fmt.Println("The string should be 6 characters")
					fmt.Println("Please enter your guess: ")
				} else {
					break
				}
			}

			// fmt.Println(text)
			totalCorrect, correctPosition = engine.StringMatch(randomStr, text)

			if totalCorrect >= targetPoint || correctPosition >= targetPoint {
				fmt.Println("Congrats..! You won the game :)")
				fmt.Println("You have guessed the correct letter but not in the correct position are: ", totalCorrect)
				fmt.Println("You have guessed the correct letter in the correct position are: ", correctPosition)
				winer = "You"
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
			fmt.Println("Now computer turn..!")
			computerGuess := random.RandomString(6)
			fmt.Println("Computer Guess: ", computerGuess)
			totalCorrect, correctPosition = engine.StringMatch(randomStr, computerGuess)
			if totalCorrect >= targetPoint || correctPosition >= targetPoint {
				fmt.Println("Oops sorry, Computer won the game :(")
				fmt.Println("Good luck for next time. Hope you enjoyed this.")
				fmt.Println("Computer guessed the correct letter but not in the correct position are: ", totalCorrect)
				fmt.Println("Computer guessed the correct letter in the correct position are: ", correctPosition)
				winer = "Computer"
				break
			} else {
				fmt.Println("Oops, even computer not guessed the correct characters this time")
				fmt.Println("Computer guessed the correct letter but not in the correct position are: ", totalCorrect)
				fmt.Println("Computer guessed the correct letter in the correct position are: ", correctPosition)
				turn = "Player"
			}
		}
	}

	return winer
}
