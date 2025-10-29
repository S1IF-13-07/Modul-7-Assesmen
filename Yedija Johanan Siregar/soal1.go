package main

import "fmt"

func main() {
  var ganjil int
  fmt.Print("Input bilangan ganjil: ")
  fmt.Scan(&ganjil)

  for i := 0; i < ganjil; i++ {
  fmt.Print(2*i + 1)
  }
  fmt.Println()
}
