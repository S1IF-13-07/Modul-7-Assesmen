package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan hari awal dan hari akhir (x y): ")
	fmt.Scan(&x, &y)

	total := 1
	for i := x; i <= y; i++ {
		total *= i
	}

	fmt.Println("Jumlah bakteri terakhir:", total)
}