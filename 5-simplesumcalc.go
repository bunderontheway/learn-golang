package main

import "fmt" // Package untuk melakukan input / output

func main(){
  
  fmt.Print("Masukkan angka pertama  :") // Menampilkan string
  var angkaPertama int // Deklarasi variabel angkaPertama sebagai integer
  fmt.Scanln(&angkaPertama) // Fungsi untuk menyimpan input yang akan disimpan di angkaPertama
  
  fmt.Print("\n Masukkan angka kedua    :") // Menampilkan string
  var angkaKedua int // Deklarasi variabel angkaKedua sebagai integer
  fmt.Scanln(&angkaKedua) // Fungsi untuk menyimpan input yang akan disimpan di angkaKedua
  
  var hasilJumlah int = angkaPertama + angkaKedua 
  // Deklarasi variabel hasilJumlah sebagai integer
  // Berisi penjumlahan dari angkaPertama dan angkaKedua yang telah diinputkan
  
  fmt.Print("\n Hasil Penjumlahan   : %i ", hasilJumlah) 
  //Menampilkan string, %i untuk menampilkan integer dari variabe hasilJumlah
  
}
