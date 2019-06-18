package main

import "fmt"

func main(){
  var score = 7
  
  switch score {
  case 10 : 
    fmt.Println("Excellent")
  case 8 :
    fmt.Println("Good Job")
  case 4 :
    fmt.Println("Try Again")
  default :
    fmt.Println("Not Bad")
  }
}
