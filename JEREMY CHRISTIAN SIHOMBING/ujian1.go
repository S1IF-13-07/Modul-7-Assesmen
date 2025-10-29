package main
import "fmt"

func main() {
    var n, i int
	fmt.Print("Masukkan angka: ")
    fmt.Scan(&n)

    fmt.Print("Output: ")
    for i = 2; i <= n*2; i += 2 {
        fmt.Print(i, " ")
    }
}