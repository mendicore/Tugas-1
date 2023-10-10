package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int

	fmt.Print("Masukkan bilangan 1-9: ")
	fmt.Scanln(&input)

	bilanganStr := strconv.Itoa(input)

	bilanganGandakan, _ := strconv.Atoi(bilanganStr + bilanganStr)

	fmt.Printf("Hasil penggandaan: %d\n", bilanganGandakan)
}
