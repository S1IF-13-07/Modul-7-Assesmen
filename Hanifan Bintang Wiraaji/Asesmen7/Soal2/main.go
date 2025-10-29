package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var c = a
	for i := ( a + 1 ); i <= b; i++ {
		c = c * i
	}
	fmt.Print(c)
}
