package main

import "fmt"

func main(){
  fmt.Print("Masukkan angka pertama  :")
  var angkaPertama int
  fmt.Scanln(&angkaPertama)
  fmt.Print("\n Masukkan angka kedua    :")
  var angkaKedua int
  fmt.Scanln(&angkaKedua)
  
  var hasilJumlah int = angkaPertama + angkaKedua
  
  fmt.Print("\n Hasil Penjumlahan   : %i ", hasilJumlah)
}
