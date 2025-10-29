package main
import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan keping : ")
	fmt.Scanln(&keping)
	
	ikat := keping / 10
	sisakeping := keping % 10 
	karung := ikat / 10 
	sisaikat := ikat % 10
	peti := karung / 8
	sisakarung := karung % 8

	fmt.Printf("Peti : %d\n", peti)
	fmt.Printf("karung : %d\n", sisakarung)
	fmt.Printf("ikat : %d\n", sisaikat)
	fmt.Printf("keping : %d\n", sisakeping)


}