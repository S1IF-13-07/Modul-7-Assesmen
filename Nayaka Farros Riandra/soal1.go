package main

import "fmt"

func main() {
	var a, ganjil int

	fmt.Print("berapa kali = ")
	fmt.Scan(&a)

	ganjil = 1
	for i := 1; i <= a; i++ {
		fmt.Print(" ", ganjil)
		ganjil += 2
	}
}