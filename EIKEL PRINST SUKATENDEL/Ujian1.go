package main
import "fmt"

func main() {
	var n int
	fmt.Print("Input n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i*2)
	}
}