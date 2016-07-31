package main

import "testing"

func TestMain(t *testing.T) {
	randomStr := "iglbrq"
	text := "iglbrg"
	totalCorrect, correctPosition := stringMatch(randomStr, text)
	if totalCorrect != 1{
		t.Error("Expected totalCorrect but not in same position 1, but got ",totalCorrect)	
	}
	if correctPosition != 5{
		t.Error("Expected totalCorrect are in same position 5, but got ",correctPosition)	
	}
}

func TestMain1(t *testing.T) {
	randomStr := "KMAFYM"
	text := "ABCDEF"
	totalCorrect, correctPosition := stringMatch(randomStr, text)
	if totalCorrect != 2{
		t.Error("Expected totalCorrect but not in same position 2, but got ",totalCorrect)	
	}
	if correctPosition != 0{
		t.Error("Expected totalCorrect are in same position 0, but got ",correctPosition)	
	}
}

// Murtuza test case
func TestMain2(t *testing.T) {
	randomStr := "AABBCC"
	text := "ABCABC"
	totalCorrect, correctPosition := stringMatch(randomStr, text)
	println("totalCorrect: ", totalCorrect)
	println("correctPosition:", correctPosition)
	if totalCorrect != 4{
		t.Error("Expected totalCorrect but not in same position 4, but got ",totalCorrect)	
	}
	if correctPosition != 2{
		t.Error("Expected totalCorrect are in same position 2, but got ",correctPosition)	
	}
}

func TestMain3(t *testing.T) {
	randomStr := "adbbcc"
	text := "abcabc"
	totalCorrect, correctPosition := stringMatch(randomStr, text)
	if totalCorrect != 4{
		t.Error("Expected totalCorrect but not in same position 4, but got ",totalCorrect)	
	}
	if correctPosition != 2{
		t.Error("Expected totalCorrect are in same position 2, but got ",correctPosition)	
	}
}