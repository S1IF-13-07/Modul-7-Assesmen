package main

import "fmt"

func main() {
	var keping, peti, karung, ikat, sisa int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	peti = keping / 800
	sisa = keping % 800

	karung = sisa / 80
	sisa = sisa % 80

	ikat = sisa / 8
	keping = sisa % 8

	fmt.Print("Hasil: ")
	fmt.Print(peti, "peti, ")
	fmt.Print(karung, "karung, ")
	fmt.Print(ikat, "ikat, dan")
	fmt.Print(sisa, "keping\n")
}