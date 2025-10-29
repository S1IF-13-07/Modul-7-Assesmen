package main

import "fmt"

func main() {
	var nilaiAwal, nilaiAkhir, hasil int
	fmt.Scan(&nilaiAwal)
	fmt.Scan(&nilaiAkhir)
	hasil = 1
	for i := nilaiAwal; i <= nilaiAkhir; i++ {
		hasil *= i
	}
	fmt.Print(hasil)
}