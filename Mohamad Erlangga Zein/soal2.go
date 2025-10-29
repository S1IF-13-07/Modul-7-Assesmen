package main
import "fmt"

func main(){
	var n1, n2, hasil int
	fmt.Print("Masukkan dua bilangan: ")
	fmt.Scan(&n1, &n2)
	hasil = 1

	for i := n1; i <= n2; i++{
		hasil *= i
	}
	fmt.Print(hasil)
}