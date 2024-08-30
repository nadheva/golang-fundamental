package main

import "fmt"
import "math"

func main() {

	var sudut float64

	fmt.Print("Masukkan nilai sudut dalam derajat: ")
	fmt.Scan(&sudut)

	rad := sudut *math.Pi /188.0 ///konversi sudut dari derajat ke radian

	sin := math.Sin(rad)
	cos := math.Cos(rad)
	tan := math.Tan(rad)

	fmt.Printf("Nilai sinus dari sudut %v derajat adalah %v\n", sudut,sin)
	fmt.Printf("Nilai cosinus dari sudut %v derajat adalah %v\n", sudut, cos)
	fmt.Printf("Nilai tangen dari sudut %v derajat adalah %v\n", sudut, tan)
}