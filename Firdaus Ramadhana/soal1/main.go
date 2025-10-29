package main 

import "fmt"

func main() {
	var j, n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	for j = 1; j <= n ; j++ {
		fmt.Print(2*j-1, " ")
	}
}