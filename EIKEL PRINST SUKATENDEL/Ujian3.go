package main

import "fmt"

func main() {
	var k int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&k)

	peti := k / 800
	k %= 800
	karung := k / 80
	k %= 80
	ikat := k / 8
	k %= 8

	fmt.Printf("%d peti, %d karung, %d ikat, %d keping\n", peti, karung, ikat, k)
}