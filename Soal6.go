package main

import "fmt"

func calculateFunctionValue(x float64) float64 {
	return (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)
}

func main() {
	var x float64

	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)

	result := calculateFunctionValue(x)

	fmt.Printf("Nilai fungsi f(x) untuk x=%.2f: %.6f\n", x, result)
}
