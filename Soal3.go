package main

import "fmt"

func main() {
	var panjang, lebar int

	fmt.Print("Masukkan panjang: ")
	fmt.Scanln(&panjang)

	fmt.Print("Masukkan lebar: ")
	fmt.Scanln(&lebar)

	keliling := (2 * panjang) + (2 * lebar)

	fmt.Printf("Keliling persegi panjang: %d\n", keliling)
}
