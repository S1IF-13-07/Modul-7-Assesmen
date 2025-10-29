package main
import "fmt"
func main() {
	var keping int
	fmt.Print("Masukan jumlah keping: ")
	fmt.Scan(&keping)

	peti := keping / 800
	sisa := keping % 800

	karung := sisa / 80
	sisa = sisa % 80

	ikat := sisa / 8
	kepingSisa := sisa % 8

	fmt.Printf("%d keping = %d peti, %d karung, %d ikat, %d keping\n", keping, peti, karung, ikat,kepingSisa)

}