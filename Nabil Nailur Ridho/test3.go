package main

import "fmt"

func main(){
	var peti, karung, ikat, keping int
	fmt.Print("Masukkan uang: ")
	fmt.Scan(&keping)

	peti = keping / 800
	karung = keping % 800 / 80
	ikat = keping % 800 % 80 / 8
	keping = keping % 8

	fmt.Print(peti," peti, ",karung," karung, ",ikat," ikat, ",keping," keping")
}