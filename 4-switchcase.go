package main

import "fmt" // Package untuk melakukan input / output

func main(){
  var score = 7 // Deklarasi variabel score
  
  switch score { // Mengambil kondisi dari variabel score
    
  case 10 : // Jika nilai dari score adalah 10
    fmt.Println("Excellent") // Maka ini yang akan ditampilkan
    
  case 8 : // Jika nilai dari score adalah 8
    fmt.Println("Good Job") //Maka ini yang akan ditampilkan
    
  case 4 : // Jika nilai dari score adalah  4
    fmt.Println("Try Again") // Maka ini yang akan ditampilkan
    
  default : // Default adalah kondisi jika nilai dari score tidak ada di salah satu point tersebut
    fmt.Println("Not Bad") // Maka ini yang akan ditampilkan
  }
}
