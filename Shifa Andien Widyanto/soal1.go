package main
import "fmt"
func main() {
	var n int
	fmt.Print("Masukan bilangan bulat n : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print(i*2-1," ")
	}
}