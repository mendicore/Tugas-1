package main

import "fmt"

func main() {
	var panjang, lebar float64

	fmt.Print("Masukkan nilai panjang: ")
	fmt.Scanln(&panjang)

	fmt.Print("Masukkan nilai lebar: ")
	fmt.Scanln(&lebar)

	luas := panjang * lebar

	fmt.Printf("Luas persegi panjang: %.2f\n", luas)
}
