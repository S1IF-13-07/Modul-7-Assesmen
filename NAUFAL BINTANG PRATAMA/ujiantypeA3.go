package main

import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	// 1 peti = 800 keping, 1 karung = 80 keping, 1 ikat = 8 keping
	peti := keping / 800
	sisa := keping - (peti * 800)

	karung := sisa / 80
	sisa = sisa - (karung * 80)

	ikat := sisa / 8
	kepingSisa := sisa - (ikat * 8)

	fmt.Printf("%d peti %d karung %d ikat %d keping\n", peti, karung, ikat, kepingSisa)
}