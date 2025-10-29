package main

import "fmt"

func main() {
    var x, y int
    
    fmt.Print("Masukkin hari awal dan hari akhir wokk: ")
    fmt.Scan(&x, &y)
    
    hasil := 1
    for i := x; i <= y; i++ {
        hasil *= i
    }
    
    fmt.Println("total bakteri ada:",hasil)
}