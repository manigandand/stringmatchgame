package engine

func StringMatch(randomStr string, text string) (totalCorrect int, correctPosition int) {
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
