package random

import (
	"math/rand"
	"strings"
	"time"
)

func ComputerRandomString(strlen int) string {
	// get current available charaters to genearte random string
	chars := CurentAvailability()
	// fmt.Println("Remaining Constant to generate Computer Guess: ", chars)

	// "Previous Guess String: "
	popin := CurrentRightGuess()
	lengthOfCorrectGuess := 0

	// generate random string
	rand.Seed(time.Now().UTC().UnixNano())
	random := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		random[i] = chars[rand.Intn(len(chars))]
	}

	for _, el := range popin {
		if el != 0 {
			// fmt.Println("PopIn: ", string(el))
			lengthOfCorrectGuess++
		}
	}
	targetPoint := GetTargetPoint()
	if lengthOfCorrectGuess == targetPoint {
		return string(popin)
	} else {
		return string(random)
	}
}

func PopoutString(popout string) {
	chars := CurentAvailability()
	// fmt.Println("Computer Guess Constant Before Popout: ", chars)
	remaining := stripchars(chars, popout)
	// fmt.Println("remaining length: ", len(remaining))
	UpdatePopoutString(remaining)
	// fmt.Println("Computer Guess Constant Before Popout: ", remaining)
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
