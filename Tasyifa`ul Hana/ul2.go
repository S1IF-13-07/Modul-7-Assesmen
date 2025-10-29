package main
import "fmt"
func main() {
	var x, y int
	fmt.Print("Masukan hari mulai (x) dan hari akhir (y): ")
	fmt.Scan(&x, &y)

	jumlah := 1
	for i := x; i <= y; i++ {
		jumlah *= i
	}
	fmt.Println("Jumlah bakteri pada hari ke-", y, "adalah:", jumlah)
}