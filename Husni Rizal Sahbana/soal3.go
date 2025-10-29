package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	peti := n / 800
	remaining := n % 800
	karung := remaining / 80
	remaining = remaining % 80
	ikat := remaining / 8
	keping := remaining % 8

	fmt.Printf("%d peti %d karung %d ikat %d keping\n", peti, karung, ikat, keping)
}
