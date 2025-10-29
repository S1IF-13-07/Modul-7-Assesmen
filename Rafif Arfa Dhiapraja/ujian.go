package main

import "fmt"

func main(){
	var bilangan int
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&bilangan)

	for i := 1; i <= bilangan; i++ {
		hasil := i * 2 - 1
		fmt.Print(hasil, " ")
	}

	
}