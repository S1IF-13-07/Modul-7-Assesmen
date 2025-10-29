package main
import "fmt"
func main() {

	var x int
	var y int
	fmt.Print("Masukan bilngan x")
	fmt.Scan(&x)
	fmt.Print("Masukan bilangan y")
	fmt.Scan(&y)

	hasil := 1
	for i := x;i <= y; i++{
		hasil= hasil * i
	}
	fmt.Print("Masukan hasil jumlah pada bakteri:", hasil)
}