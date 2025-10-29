package main 

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Scan(&a, &b)
	hasil = 1
	for i := a; i <= b ; i++ {
		hasil *= i
	}
	fmt.Print(hasil)
}