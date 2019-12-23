package main

import "fmt" // package untuk melakukan input / output

func main() {
  var nilai = 80 // Deklarasi variabel
  
  if nilai == 100{ //Jika variabel nilai yang dideklarasikan == 100
    fmt.Println("Nilai anda A")  // Maka ini yang akan ditampilkan
    
  }else if nilai > 75{ //Jika variabel nilai yang dideklarasikan > 75
    fmt.Println("Nilai anda B") // Maka ini yang akan ditampilkan
    
  }else if nilai >= 50 && nilai <= 75 { // Jika variabel nilai yang dideklarasikan >= 50 dan <= 75
    fmt.Println("Nilai anda C") // Maka ini yang akan ditampilkan
    
  }else { // Jika variabel nilai tidak berisi satupun dari ketiga poin diatas
    fmt.Println("Anda harus mengulang") // Maka ini yang akan ditampilkan
  }
  
}
