package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	hasil := 1
	for i := x; i <= y; i++ {
		hasil *= i
	}

	fmt.Println("Output:", hasil)
}
