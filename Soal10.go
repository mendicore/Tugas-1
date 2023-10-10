package main

import "fmt"

func main() {
	var panjang, lebar, tinggi, berat int

	fmt.Print("Masukkan panjang paket (cm): ")
	fmt.Scanln(&panjang)

	fmt.Print("Masukkan lebar paket (cm): ")
	fmt.Scanln(&lebar)

	fmt.Print("Masukkan tinggi paket (cm): ")
	fmt.Scanln(&tinggi)

	fmt.Print("Masukkan berat paket (gram): ")
	fmt.Scanln(&berat)

	volumeM3 := float64(panjang*lebar*tinggi) / 1000000.

	bolehDikirim := volumeM3 <= 1.0 && berat <= 30000

	fmt.Println("Apakah paket boleh dikirimkan? ", bolehDikirim)
}
