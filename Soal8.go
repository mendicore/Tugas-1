package main

import "fmt"

func main() {
	var b1, b2, b3 int

	fmt.Print("Masukkan berat belerang b1 (gram): ")
	fmt.Scanln(&b1)

	fmt.Print("Masukkan berat belerang b2 (gram): ")
	fmt.Scanln(&b2)

	fmt.Print("Masukkan berat belerang b3 (gram): ")
	fmt.Scanln(&b3)

	totalBeratGram := b1 + b2 + b3

	beratKg := totalBeratGram / 1000
	sisaGram := totalBeratGram % 1000

	fmt.Printf("Berat belerang: %d kilogram %d gram\n", beratKg, sisaGram)
}
