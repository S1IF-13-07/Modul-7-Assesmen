package main

import "fmt"

func main() {
	var peti, karung, ikat, keping int

	fmt.Print("masukkan jumlah keping = ")
	fmt.Scan(&keping)

	peti = keping / 800
	karung = keping % 800 / 100
	ikat = keping % 800 % 100 / 10
	keping = keping % 800 % 10

	fmt.Println("peti =", peti)
	fmt.Println("karung =", karung)
	fmt.Println("ikat =", ikat)
	fmt.Println("keping =", keping)  
}