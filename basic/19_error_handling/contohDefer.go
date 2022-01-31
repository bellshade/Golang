package main

import "fmt"

// membuat fungsi pertambahan
func tambah(angka1, angka2 int) int {
  hasil := angka1 + angka2
  fmt.Println("hasil pertambahan :", hasil)

  return 0
}

func main() {
  fmt.Println("memulai")

  // membuat statement defer
  // ekseskusi dengan fungsi LIFO
  defer fmt.Println("akhir")
  defer tambah(33, 20)
  defer tambah(20, 20)
}

