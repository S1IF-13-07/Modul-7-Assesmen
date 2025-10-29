package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)
	jumlah := 1
	for i := x; i <= y; i++ {
		jumlah *= i
	}
	fmt.Println(jumlah)
}