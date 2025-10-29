package main
import "fmt"

func main() {
    var x, y int
	fmt.Print("Masukkan hari awal dan hari akhir: ")
    fmt.Scan(&x, &y)

	hasil := 1
    for i := x; i <= y; i++ {
        hasil *= i
    }
	fmt.Println("Jumlah bakteri terakhir dari hari x sampai hari y:", hasil)
}