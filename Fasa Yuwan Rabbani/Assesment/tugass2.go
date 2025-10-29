package main
import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Masukkan hari pertama : ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan hari ke n : ")
	fmt.Scanln(&b)

	jumlah := 1
	for i := a; i <= b; i++ {
		jumlah *= i
	}
	hasil := jumlah
	fmt.Print(hasil)
}