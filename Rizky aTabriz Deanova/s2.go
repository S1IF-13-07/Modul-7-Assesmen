package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x > y {
		x, y = y, x
	}
	bakteri := 1
	for i := x; i <= y; i++ {
		bakteri = bakteri * i
	}
	fmt.Println(bakteri)
}
