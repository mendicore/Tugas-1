package main

import "fmt"

func main() {
	var bilanganAdik, bilanganKakak int

	fmt.Print("Masukkan bilangan tebakan adik: ")
	fmt.Scanln(&bilanganAdik)

	fmt.Print("Masukkan bilangan tebakan kakak: ")
	fmt.Scanln(&bilanganKakak)

	adikMenang := (bilanganAdik%2 == 1 && bilanganKakak%2 == 0) || (bilanganAdik%2 == 0 && bilanganKakak%2 == 1)

	fmt.Println("Apakah adik pemenang? ", adikMenang)
}
