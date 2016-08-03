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

	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
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
