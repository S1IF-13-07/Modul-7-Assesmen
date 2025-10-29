package main
import "fmt"
func main() {
	var n int
	fmt.Print("Masukan jumlah bilangan genap: ")
	fmt.Scan(&n)

	fmt.Print("Deret bilangan genap: ")
	for i := 1; i <= n; i++ {
	fmt.Print(2*i, " ")
	}
}
