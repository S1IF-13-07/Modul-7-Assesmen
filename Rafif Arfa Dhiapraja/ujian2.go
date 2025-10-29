package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Hari pertama bakteri: ")
	fmt.Scan(&x)
	fmt.Print("Hari Terakhir bakteri: ")
	fmt.Scan(&y)

	hasil := 1

	for i := x; i <= y; i++ {
		hasil *= i
	}
	
	fmt.Print("Jumlah Bakteri: ", hasil)
}
