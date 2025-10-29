package main
import "fmt"
func main() {
	var peti, karung, ikat, keping int
	fmt.Print("Masukan uang dalan satuan keping : ")
	fmt.Scan(&keping)
	
	peti = keping / 800
	karung = (keping % 800) / 80
	ikat = (keping % 80) / 8
	keping = keping % 8

	fmt.Println(peti, "peti", karung, "karung", ikat, "ikat", keping, "keping")
}
