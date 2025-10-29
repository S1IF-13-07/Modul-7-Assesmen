package main

import "fmt"

func main() {
  var uang int
  fmt.Print("Masukkan jumlah uang: ")
  fmt.Scan(&uang)

peti := uang / 800
sisa := uang % 800

karung := sisa / 100
sisa = sisa % 100

ikat := sisa / 10
keping := sisa % 10

  fmt.Printf("%d %d %d %d\n", peti, karung, ikat, keping)

  fmt.Println("Hasil dalam bentuk peti: ", peti)
  fmt.Println("Hasil dalam bentuk karung: ", karung)
  fmt.Println("Hasil dalam bentuk ikat: ", ikat)
  fmt.Println("Hasil dalam bentuk keping: ", keping)
}
