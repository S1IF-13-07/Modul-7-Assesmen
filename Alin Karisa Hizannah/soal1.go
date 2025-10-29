package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan Bilangan n : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print(2*i, " ")
	}
}
