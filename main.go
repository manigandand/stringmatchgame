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
  var lastCorrectChar byte
  for i := 0; i < len(byterandomStr); i++ {
    // fmt.Println(userInput[i]," ",byterandomStr[i])
    if byterandomStr[i] == userInput[i] {
      correctPosition++
      /*
        Calculate the strings gused correctly but not in the same position
      */
      for j := i+1; j < len(byterandomStr); j++{
        if lastCorrectChar != byterandomStr[i]{
          if byterandomStr[i] == userInput[j] {
            // fmt.Println("r: ",i," ",byterandomStr[i]," u: ",j," ",userInput[j])     
            if byterandomStr[j] != userInput[j] {
              totalCorrect++
            }
          }
        }
      }
      lastCorrectChar  = byterandomStr[i]

    } else{
      /*
        Calculate the strings gused correctly but not in the same position
      */
      for j := 0; j < len(byterandomStr); j++{
        if lastCorrectChar != byterandomStr[i]{
          if byterandomStr[i] == userInput[j] {
            // fmt.Println("r: ",i," ",byterandomStr[i]," u: ",j," ",userInput[j])  
            if byterandomStr[j] != userInput[j] {
              totalCorrect++
            }
          }
        }
      }
      lastCorrectChar  = byterandomStr[i]
    }
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
