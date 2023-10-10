package main

import "fmt"

func main() {
	var suhuCelsius float64

	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scanln(&suhuCelsius)

	suhuKelvin := suhuCelsius + 273.15

	fmt.Printf("Suhu dalam Kelvin: %.2f\n", suhuKelvin)
}
