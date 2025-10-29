package main
import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	kepingPerIkat := 8
	ikatPerKarung := 10
	karungPerPeti := 10

	kepingPerKarung := kepingPerIkat * ikatPerKarung
	kepingPerPeti := kepingPerKarung * karungPerPeti

	peti := keping / kepingPerPeti
	sisa := keping % kepingPerPeti

	karung := sisa / kepingPerKarung
	sisa %= kepingPerKarung

	ikat := sisa / kepingPerIkat
	sisa %= kepingPerIkat

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, sisa)
}
