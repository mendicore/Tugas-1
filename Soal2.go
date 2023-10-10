package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64
	const pi float64 = 3.14

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scanln(&jariJari)

	volume := (4.0 / 3.0) * pi * math.Pow(jariJari, 3)

	fmt.Printf("Volume bola: %.2f\n", volume)
}
