package random

import (
	"io"
	"io/ioutil"
	"os"
)

const path = "src/stringMatchGame/random/data/computermove.txt"

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
	input := []byte("abcdefghijklmnopqrstuvwxyz")

	err := ioutil.WriteFile(path, input, 0644)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
