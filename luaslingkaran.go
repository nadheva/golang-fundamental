package main

import "fmt"

func main() {
	var jari float64

	fmt.Print("Masukkan panjang jari-jari lingkaran: ")
	fmt.Scan(&jari)

	luas := 22/7 * jari*jari
	fmt.Println("Luas lingkarah adalah: ", luas)
}