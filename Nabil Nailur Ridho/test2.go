package main

import "fmt"

func main(){
	var x, y, hasil int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)

	hasil = 1
	for i := x; i <= y; i++ {
		hasil *= i
	}
	fmt.Print(hasil)
}