package main

import "fmt"

func main() {
  var nilai = 80
  
  if nilai == 100{
    fmt.Println("Nilai anda A")  
  }else if nilai > 75{
    fmt.Println("Nilai anda B")
  }else if nilai >= 50 && nilai <= 75 {
    fmt.Println("Nilai anda C")
  }else {
    fmt.Println("Anda harus mengulang")
  }
  
}
