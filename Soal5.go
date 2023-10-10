package main

import "fmt"

func main() {
	var x int
	fmt.Print("masukan nilai x: ")
	fmt.Scanln(&x)

	var y = (3*x - 5) * (2*x + 1)

	fmt.Print(y)
}
