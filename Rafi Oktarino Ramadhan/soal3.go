package main

import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	
	kepingPerIkat := 10
	ikatPerKarung := 10
	karungPerPeti := 8

	
	kepingPerPeti := karungPerPeti * ikatPerKarung * kepingPerIkat

	
	peti := keping / kepingPerPeti
	sisa := keping % kepingPerPeti

	
	kepingPerKarung := ikatPerKarung * kepingPerIkat
	karung := sisa / kepingPerKarung
	sisa = sisa % kepingPerKarung

	
	ikat := sisa / kepingPerIkat

	
	sisaKeping := sisa % kepingPerIkat

	
	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, sisaKeping)
}
