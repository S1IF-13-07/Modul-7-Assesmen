package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fmt.Println("Deret bilangan genap: ")
	for i :=1; i <= n; i ++ {
		fmt.Print(i*2," ")
	}
}