package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan hari mulai (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan hari akhir (y): ")
	fmt.Scan(&y)

	jumlah := 1

	for i := 1; i < x; i++ {
		jumlah *= i
	}

	for i := x; i <= y; i++ {
		jumlah *= i
	}

	fmt.Println("Jumlah bakteri pada hari", y, "adalah:", jumlah)
}
