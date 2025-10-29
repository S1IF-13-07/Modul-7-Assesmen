package main

import "fmt"

func main() {
	var keping, ikat, karung, peti int
	fmt.Scan(&keping)
	peti = keping / 800
	karung = keping % 800 / 100
	ikat = keping % 800 % 100 / 10
	keping = keping % 10
	fmt.Printf("%d %d %d %d", peti, karung, ikat, keping)
}
