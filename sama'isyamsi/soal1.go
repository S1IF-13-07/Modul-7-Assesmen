package main
import "fmt"
func main() {

	var n int
	var hasil int
	fmt.Print("Masukan bilangan n")
	fmt.Scan(&n)
	fmt.Println("Deret bilangan ganjil")
	for i := 1; i <= n; i++ {
		fmt.Print(i*2-1," ")

	}
	fmt.Println("Hasil dari 1 sampai", n, "adalah:", hasil)
}
