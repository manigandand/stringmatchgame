package random

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var path = "random/data/computermove.txt"
var path1 = "random/data/computer_correct_moves.txt"
var path2 = "random/data/targetpoint.txt"

func init() {
	var err error
	path, err = filepath.Abs(path)
	CheckError(err)
	path1, err = filepath.Abs(path1)
	CheckError(err)
	path2, err = filepath.Abs(path2)
	CheckError(err)
}

// const flag = true

// CurentAvailability ...
func CurentAvailability() string {
	file, err := os.Open(path)
	CheckError(err)

	// close on last
	defer func() {
		err := file.Close()
		CheckError(err)
	}()

	// make a buffer to keep chunks that are read
	buffer := make([]byte, 26)
	m := 0
	for {
		// read a chunk
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		m = n
	}

	// string()
	return string(buffer[:m])
}

func UpdatePopoutString(update string) {
	input := []byte(update)

	err := ioutil.WriteFile(path, input, 0644)
	CheckError(err)
}

// RefillInputFile ...
func RefillInputFile() {
	// regenerate string to random generate
	str := "abcdefghijklmnopqrstuvwxyz"
	input := []byte(str)

	err := ioutil.WriteFile(path, input, 0644)
	CheckError(err)

	// reset computer right guess file
	input1 := []byte("")
	err = ioutil.WriteFile(path1, input1, 0644)
	CheckError(err)

	// reset target point file
	input2 := []byte("")
	err = ioutil.WriteFile(path2, input2, 0644)
	CheckError(err)
}

/*
*
*~~~~~~~~~~~ SAVE ALL THE CORRECT GUESS CHARCTERS ~~~~~~~~~~~
*
 */
func CurrentRightGuess() []byte {
	file, err := os.Open(path1)
	CheckError(err)

	// close on last
	defer func() {
		err := file.Close()
		CheckError(err)
	}()

	// make a buffer to keep chunks that are read
	buffer := make([]byte, 6)
	for {
		// read a chunk
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}

	// string()
	return buffer
}

func UpdatePopinString(popin string) {
	input := []byte(popin)

	err := ioutil.WriteFile(path1, input, 0644)
	CheckError(err)
}

/*
*
*~~~~~~~~~~~ SAVE ALL THE CORRECT GUESS CHARCTERS ~~~~~~~~~~~
*
 */
func GetTargetPoint() int {
	file, err := os.Open(path2)
	CheckError(err)

	// close on last
	defer func() {
		err := file.Close()
		CheckError(err)
	}()

	// make a buffer to keep chunks that are read
	buffer := make([]byte, 1)
	for {
		// read a chunk
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}

	x, _ := strconv.Atoi(string(buffer))
	return x
}

func UpdateTargetPoint(tp int) {
	x := strconv.Itoa(tp)

	input := []byte(x)

	err := ioutil.WriteFile(path2, input, 0644)
	CheckError(err)
}

// Error Handler
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
