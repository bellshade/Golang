package main

import "fmt"

func main(){
  var angka int = 20
  
  // membuat pointer
  var poin *int

  // menyimpan address dari angka ke variabel pointer
  poin = &angka

  fmt.Printf("alamat dari variabel angka: %x\n", &angka)
  
  // alamat yang kesimpan dalam variabel pointer
  fmt.Printf("alamat dari variabel pointer: %x\n", poin)

  // mengakses value dari pointer
  fmt.Printf("value dari variabel poin adalah: %d\n", *poin)

  // menambahkan value yang terdapat pada variabel poin
  var angka_hasil = *poin + 20
  fmt.Printf("hasil value variabel poin sesudah ditambahkan: %d\n", angka_hasil)  
}

