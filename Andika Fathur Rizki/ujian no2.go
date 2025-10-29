package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("masukan jumlah bakteri hari pertama : ")
	fmt.Scan(&x)
	fmt.Print("masukan jumlah bakteri hari terakhir : ")
	fmt.Scan(&y)

	jumlahbakteri := 1

	for i := x; i <= y; i++ {
		jumlahbakteri *= i
	}

	fmt.Print("jumlah bakteri dari hari pertama sampai terakhir adalah : ", jumlahbakteri)

}
