package engine

import "stringMatchGame/random"

func StringMatch(randomStr string, text string, turn string) (totalCorrect int, correctPosition int, popout string) {
	byterandomStr := []byte(randomStr)
	userInput := []byte(text)

	totalCorrect = 0
	correctPosition = 0
	/*
	   Check for the given string matches exactly in the same position
	*/
	lastCheckedChar := make([]string, 6)
	// "Previous Guess String: "
	popin := random.CurrentRightGuess()
	popout = ""

	for i := 0; i < len(byterandomStr); i++ {
		// fmt.Println(userInput[i]," ",byterandomStr[i])
		if byterandomStr[i] == userInput[i] {
			correctPosition++
			popin[i] = userInput[i]
		}
		/*
		   Calculate the strings guesed correctly but not in the same position
		*/
		for j := 0; j < len(userInput); j++ {
			// check wheather the current index already searched or not
			if checkIndex(lastCheckedChar, byterandomStr[i]) {
				if byterandomStr[i] == userInput[j] {
					// fmt.Println("r: ", i, " ", string(byterandomStr[i]), " u: ", j, " ", string(userInput[j]))
					if j >= i {
						if byterandomStr[j] != userInput[j] {
							totalCorrect++
							popin[i] = userInput[j]
						}
					} else {
						totalCorrect++
						popin[i] = userInput[j]
					}
				}
			}
		}
		lastCheckedChar[i] = string(byterandomStr[i])

	}

	// find wrong guess to popout
	if turn == "Computer" {
		// for i := 0; i < len(userInput); i++ {
		// 	for j := 0; j < len(byterandomStr); j++ {
		// 		if userInput[i] != byterandomStr[j] {
		// 			if j == 5 {
		// 				popout += string(userInput[i])
		// 				// fmt.Println("u: ", i, " ", string(userInput[i]), " r: ", j, " ", string(byterandomStr[j]))
		// 			}
		// 		} else {
		// 			break
		// 		}
		// 	}
		// }

		// popout all the characteres computer guessed this time
		// so the next time the correct character won't repeat
		popout = string(userInput)

		// make byte slice of all correct characteres to save
		// fmt.Println("popin: ", popin)
		// fmt.Println("popin: ", string(popin))
		random.UpdatePopinString(string(popin))
	}

	// make byte of all popout characteres
	popoutRet := []byte(popout)

	return totalCorrect, correctPosition, string(popoutRet)
}
