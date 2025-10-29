package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan genap: ")
	fmt.Scan(&n)

	fmt.Print("Deret bilangan genap: ")
	for i := 1; i <= n; i++ {
		fmt.Print(2*1, " ")
	}
}