// package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	// open input file

// 	d1 := []byte("abcdefghijklmnopqrstuvwxyz")
// 	err := ioutil.WriteFile("input.txt", d1, 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fi, err := os.Open("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// close fi on exit and check for its returned error
// 	defer func() {
// 		if err := fi.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	// open output file
// 	fo, err := os.Create("output.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// close fo on exit and check for its returned error
// 	defer func() {
// 		err := os.Remove("output.txt")
// 		if err = fo.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	// make a buffer to keep chunks that are read
// 	buf := make([]byte, 1024)
// 	for {
// 		// read a chunk
// 		n, err := fi.Read(buf)
// 		if err != nil && err != io.EOF {
// 			panic(err)
// 		}
// 		fmt.Println(n)
// 		fmt.Println(string(buf))
// 		if n == 0 {
// 			break
// 		}

// 		// write a chunk
// 		if _, err := fo.Write(buf[:n]); err != nil {
// 			panic(err)
// 		}

// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const s = "abcdefghijklmnopqrstuvwxyz"

// func stripchars(str, chr string) string {
// 	return strings.Map(func(r rune) rune {
// 		if strings.IndexRune(chr, r) < 0 {
// 			return r
// 			fmt.Println(r)
// 		}
// 		return -1
// 	}, str)
// }

// func main() {
// 	fmt.Println("The Orginal String: ", s)
// 	fmt.Println()
// 	fmt.Println()
// 	reader := bufio.NewReader(os.Stdin)
// 	i := 2
// 	for i < 10 {
// 		if i%2 == 0 {
// 			fmt.Println("Enter Character to remove: ")
// 			text, _ := reader.ReadString('\n')
// 			result := stripchars(s, text)
// 			fmt.Println(result)
// 			fmt.Println()
// 		} else {
// 			fmt.Println("not you: ")
// 		}
// 		fmt.Println(i)
// 		i++
// 	}
// }

// package main

// import "fmt"
// import "math/rand"
// import "time"

// func RandomString(strlen int) string {
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	const chars = "abcdefghijklmnopqrstuvwxyz"
// 	result := make([]byte, strlen)
// 	for i := 0; i < strlen; i++ {
// 		fmt.Println(rand.Intn(len(chars)))
// 		//fmt.Println(chars[rand.Intn(len(chars))])
// 		result[i] = chars[rand.Intn(len(chars))]
// 	}
// 	return string(result)
// }

// func main() {
// 	fmt.Println(RandomString(6))
// 	//fmt.Println(RandomString(20))
// }

package main

import "fmt"

func in_array(val string, array []string) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if val == v {
			index = i
			exists = true
			return
		}
	}

	return
}

func main() {
	names := []string{"Mary", "Anna", "Beth", "Johnny", "Beth"}

	fmt.Println(in_array("Jack", names))
}
