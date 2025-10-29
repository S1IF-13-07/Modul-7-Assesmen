package main
import "fmt"

func main() {
    var keping int
    fmt.Print("Masukkan jumlah keping: ")
    fmt.Scan(&keping)

    peti := keping / 800
    keping %= 800
    karung := keping / 80
    keping %= 80
    ikat := keping / 8
    keping %= 8

    fmt.Println("Keluaran konversi:")
    fmt.Println("Peti:", peti)
    fmt.Println("Karung:", karung)
    fmt.Println("Ikat:", ikat)
    fmt.Println("Keping:", keping)
}