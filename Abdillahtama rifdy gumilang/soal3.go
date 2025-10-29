package main

import "fmt"

func main() {
	var totalKeping int

	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scanln(&totalKeping)

	satuIkat := 8
	satuKarung := 10 * satuIkat
	satuPeti := 10 * satuKarung

	peti := totalKeping / satuPeti
	sisa := totalKeping % satuPeti

	karung := sisa / satuKarung
	sisa = sisa % satuKarung

	ikat := sisa / satuIkat
	keping := sisa % satuIkat

	fmt.Println(peti, "peti,", karung, "karung,", ikat, "ikat, dan", keping, "keping")
}
