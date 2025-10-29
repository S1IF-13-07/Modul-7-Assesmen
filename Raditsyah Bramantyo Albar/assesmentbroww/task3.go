package main

import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	ikat := keping / 10
	sisaKeping := keping % 10

	karung := ikat / 10
	sisaIkat := ikat % 10

	peti := karung / 8
	sisaKarung := karung % 8

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, sisaKarung, sisaIkat, sisaKeping)
}