package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan hari mulai (X) dan hari akhir (Y): ")
	fmt.Scan(&x, &y)

	jumlah := 1
	for i := x; i <= y; i++ {
		jumlah *= i
	}
	fmt.Println("Jumlah bakteri pada hari ke-", y, "adalah: ", jumlah)

}