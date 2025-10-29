package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y (pisahkan dengan spasi): ")
	fmt.Scan(&x, &y)

	Output := 1
	for i := x; i <= y; i++ {
		Output *= i
	}

	fmt.Println("Output:", Output)
}
