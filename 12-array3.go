package main

import "fmt" // Package untuk melakukan input / output

func main{

// Dekalarasi variabel array tanpa jumlah elemen
// Angka elemen cukup diganti dengan titik - titik tiga biji "..."
// Maka jumlah elemen akan mengikuti data yang diisikan
var angka = [...]int{1, 2, 3, 4, 5}


fmt.Println("data array \t:", angka) // Menampilkan isi dari elemen array
fmt.Println("jumlah elemen \t:", len(angka)) //Menampilkan jumlah elemen array dengan perintah "len"

}
