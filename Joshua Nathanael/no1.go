package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan: ")
	fmt.Scan(&n)

	fmt.Print("Pengeluaraannya Adalah: ")

	for i := 1; i <= n*2; i += 2 {
		fmt.Print(i, " ")
	}
}
