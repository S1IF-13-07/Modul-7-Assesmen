package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	hasil := 1
	for i := x; i <= y; i ++ {
		hasil = hasil * i
	}
	fmt.Println("Jumlah bakteri pada hari terakhir:", hasil)
}