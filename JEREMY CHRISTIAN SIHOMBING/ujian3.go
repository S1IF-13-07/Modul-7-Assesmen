package main
import (
	"fmt"
)

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scanln(&keping)
	peti := keping / 800
	sisa := keping % 800
	karung := sisa / 80
	sisa = sisa % 80
	ikat := sisa / 8
	sisa = sisa % 8
	fmt.Println(peti, "peti,", karung, "karung,", ikat, "ikat,", "dan", sisa, "keping")
}
