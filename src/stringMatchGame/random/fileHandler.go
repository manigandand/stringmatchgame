package random

import (
	"io"
	"io/ioutil"
	"os"
)

const path = "src/stringMatchGame/random/data/computermove.txt"
const path1 = "src/stringMatchGame/random/data/computerCorrectMove.txt"

// const flag = true

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

func RefillInputFile() {
	str := "abcdefghijklmnopqrstuvwxyz"
	input := []byte(str)

	err := ioutil.WriteFile(path, input, 0644)
	CheckError(err)

	input1 := []byte("")
	err = ioutil.WriteFile(path1, input1, 0644)
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

// Error Handler
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
