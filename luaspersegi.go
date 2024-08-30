package main

import "fmt"

func main() {
 var sisi float64

 fmt.Print("Masukkan panjang sisi persegi: ")
 fmt.Scan(&sisi)

 luas := sisi * sisi
  fmt.Println("Luas persegi adalah: ", luas)
}