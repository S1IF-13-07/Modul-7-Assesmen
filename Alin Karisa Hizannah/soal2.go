package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Print("Masukkan Hari ke x - y: ")
	fmt.Scan(&x, &y)

	hasil = 1
	for i := x; i <= y; i++ {
		hasil *= i
	}
	fmt.Println("Jumlah Bakteri terakhir:", hasil)
}
