package main

import "fmt"

func main() {
	var a int
	var hasil = -1
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		hasil += 2
		fmt.Print(hasil, " ")
	}
}
