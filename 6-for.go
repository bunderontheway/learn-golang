package main

import "fmt"

func main(){
  
  //cara 1
  for i := 0 ; i < 5 ; i++{
    fmt.Println("Angka" , i)
  }
  
  //cara 2
  var j = 0
  
  for j < 10{
    fmt.Println("Angka" , j)
    j++
  }
}
