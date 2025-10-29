package main

import (
	"fmt"
)

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	kepingPerIkat := 10
	ikatPerKarung := 10
	karungPerPeti := 8

	kepingPerKarung := kepingPerIkat * ikatPerKarung
	kepingPerPeti := kepingPerKarung * karungPerPeti

	peti := keping / kepingPerPeti
	sisa := keping % kepingPerPeti

	karung := sisa / kepingPerKarung
	sisa %= kepingPerKarung

	ikat := sisa / kepingPerIkat
	kepingSisa := sisa % kepingPerIkat

	fmt.Printf("%d Peti %d Karung %d Ikat %d Keping\n", peti, karung, ikat, kepingSisa)
}
