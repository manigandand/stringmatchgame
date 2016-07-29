package main
import (
"fmt"
"math/rand"
"time"
"bufio"
"os"
)

func main() {
  randomStr := RandomString(6)
  fmt.Println(randomStr)

  // get user input
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Please enter a string only six chars long!")
  fmt.Print("Enter text to guss: ")
  text, _ := reader.ReadString('\n')
  // fmt.Println(text)
  totalCorrect, correctPosition := stringMatch(randomStr, text)

  fmt.Println("You have gussed the correct letter but not in the correct position are: ", totalCorrect)
  fmt.Println("You have gussed the correct letter in the correct position are: ", correctPosition)  
}

func stringMatch(randomStr string, text string)(totalCorrect int, correctPosition int){  
  byterandomStr := []byte(randomStr)
  userInput := []byte(text)

  totalCorrect = 0
  correctPosition = 0  
  /*
    Check for the given string matches exactly in the same position
  */
  for i := 0; i < len(byterandomStr); i++ {
    // fmt.Println(userInput[i]," ",byterandomStr[i])
    if byterandomStr[i] == userInput[i] {
      correctPosition++
    }
  }
  /*
    Calculate the strings gused correctly but not in the same position
  */
  r := 0
  u := 0
  for _,rndStr := range byterandomStr {
    for _,usrStr := range userInput {
      if rndStr == usrStr && r != u{
        // fmt.Println("Pos: ",r," ",string(rndStr)," --> ","Pos: ",u," ",string(usrStr))
        // fmt.Println(string(rndStr), " This pressents")
        totalCorrect++
      }
      u++
    }
    r++
    u = 0
  }

  return totalCorrect, correctPosition
}
func RandomString(strlen int) string {
  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "abcdefghijklmnopqrstuvwxyz"
  result := make([]byte, strlen)
  for i := 0; i < strlen; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}
