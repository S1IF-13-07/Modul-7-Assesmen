package main
import "fmt"
func main() {
	var x, y int
	fmt.Print("Masukan nilai x :")
	fmt.Scan(&x)
	fmt.Print("Masukan nilai y : ")
	fmt.Scan(&y)

	hasil := 1
	for i := x; i <= y; i++ {
		hasil = hasil * i
	}
	fmt.Print(hasil)
}