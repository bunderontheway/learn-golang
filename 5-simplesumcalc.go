package main

import "fmt"

func main(){
  fmt.Print("Masukkan angka pertama  :")
  var angkaPertama int
  fmt.Scanf(&angkaPertama)
  fmt.Print("\n Masukkan angka kedua    :")
  var angkaKedua int
  fmt.Scanf(&angkaKedua)
  
  var hasilJumlah int = angkaPertama + angkaKedua
  
  fmt.Print("\n Hasil Penjumlahan   : %i ", hasilJumlah)
}
