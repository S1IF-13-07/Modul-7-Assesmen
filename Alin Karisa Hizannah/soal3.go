package main

import "fmt"

func main() {
	var peti, karung, ikat, keping int
	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&keping)

	peti = keping / 800
	karung = keping % 800 / 80
	ikat = keping % 800 % 80 / 8
	keping = keping % 8

	fmt.Printf("%d peti, %d karung, %d ikat, %d keping", peti, karung, ikat, keping)
}
