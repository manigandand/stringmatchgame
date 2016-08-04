package random

import (
	"fmt"
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
			fmt.Println("PopIn: ", string(el))
			lengthOfCorrectGuess++
		}
	}
	fmt.Println("lengthOfCorrectGuess: ", lengthOfCorrectGuess)
	if lengthOfCorrectGuess == 6 {
		return string(popin)
	} else {

	}

	fmt.Println("PopIn: ", string(popin))

	fmt.Println("random: ", string(random))
	// result := make([]byte, 6)

	return string(random)
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
			fmt.Println(r)
		}
		return -1
	}, str)
}
