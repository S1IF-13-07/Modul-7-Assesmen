package main
import "fmt"

func main(){
	var n, hasil int
	fmt.Print("Masukkan perulangan: ")
	fmt.Scan(&n)
	hasil = 0

	for i := 1; i <= n; i++{
		hasil += 2
		fmt.Print(hasil, " ")
	}
}