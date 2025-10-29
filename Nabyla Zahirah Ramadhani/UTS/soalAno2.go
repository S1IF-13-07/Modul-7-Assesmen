package main
import "fmt"
func main (){
	var x, y int
	fmt.Print("bilangan x, y: ")
	fmt.Scan(&x, &y)
	hasil := 1
	for i := x; i <= y; i++ {
		hasil *= i
	}

	fmt.Println("Bakteri hari akhir:", hasil)
}