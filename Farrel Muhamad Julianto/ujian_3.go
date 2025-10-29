package main

import "fmt"

func main() {
	var k int
	fmt.Print("Masukkan keping: ")
	fmt.Scan(&k)

	peti := k / 800
	karung := (k % 800) / 80
	ikat := (k % 80) / 8
	keping := k % 8

	fmt.Println(peti, "peti,", karung, "karung,", ikat, "ikat, dan", keping, "keping")
}
