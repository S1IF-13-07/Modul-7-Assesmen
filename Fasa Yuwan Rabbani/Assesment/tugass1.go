package main
import "fmt"

func main() {
	var a int
	fmt.Print("Masukkan angka bilangan bulat : ")
	fmt.Scanln(&a)

		ganjil := 0
	for i := 1; ganjil < a; i += 2 {
		fmt.Print(i, " ")
		ganjil++
	}

}