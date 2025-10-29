package main

import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	peti := keping / 800
	karung := keping % 800 / 80
	ikat := keping % 80 / 8
	keping = keping % 8

	fmt.Printf("Peti: %d, Karung: %d, Ikat: %d, Keping: %d\n", peti, karung, ikat, keping)
}