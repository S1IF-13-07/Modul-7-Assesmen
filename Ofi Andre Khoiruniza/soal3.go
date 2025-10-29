package main

import "fmt"

func main () {
	var peti, karung, ikat, keping, total int

	fmt.Scan(&total)
	peti = total / 800
	karung = (total % 800) / 100
	ikat = (total % 800 % 100) / 10
	keping = (total % 800) % 10

	fmt.Println(peti, "peti", karung, "karung", ikat, "ikat", keping, "keping")


}