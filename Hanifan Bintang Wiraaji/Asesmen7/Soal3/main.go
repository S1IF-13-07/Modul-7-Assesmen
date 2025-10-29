package main

import "fmt"

func main() {
	var keping, peti, karung, ikat int
	fmt.Scan(&keping)
	peti = keping / 800
	karung = keping % 800 / 100
	ikat = keping % 800 % 100 / 10
	keping = keping % 800 % 100 % 10
	fmt.Printf("%v peti, %v karung, %v ikat, %v keping", peti, karung, ikat, keping)
}
