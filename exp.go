package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	strlen := 6
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	fmt.Println(string(result))
}
