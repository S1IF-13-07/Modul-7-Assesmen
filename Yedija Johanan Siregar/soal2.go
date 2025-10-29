package main

import "fmt"

func main() {
  var x, y int
  fmt.Print("Masukkan hari x dan y: ")
  fmt.Scan(&x, &y)

  if x > y {
  fmt.Println("Hari x harus kurang dari atau sama dengan y")
    }

  bakteri := 1
  for i := x; i <= y; i++ {
    if i >= x {
      bakteri *= i
    }
  }

  fmt.Printf("Jumlah bakteri awal adalah: %d\n", bakteri)
}
