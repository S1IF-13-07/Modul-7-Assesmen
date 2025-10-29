package main 

import "fmt"

func main() {
	var peti, karung, ikat, keping int
	fmt.Println("Masukkan jumlah keping: ")
	fmt.Scan(&keping)
	peti = keping / 800
	karung = keping % 800 / 100
	ikat = keping % 100 / 10
	keping = keping % 10
	fmt.Printf("peti: %d, karung: %d, ikat: %d, keping: %d\n",peti , karung, ikat, keping)
}