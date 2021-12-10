package main

import "fmt"

func main(){
  // membuat array dengan jumlah elemen = 15
  var n [15] int
  var i, j int

  // inisialiasi elemen array n ke 0
  for i = 0; i < 15; i++{
    // set elemen pada lokasi i = i + 3
    // jadi memulai pada elemen 0 yaitu 3
    // 3,4,5,6,7,8,9,10....,17
    // sampai elemen ke 15
    n[i] = i + 3
  }
  
  // menampilkan hasil elemen pada array
  for j = 0; j < 15; j++{
    fmt.Printf("elemen[%d] = %d\n", j, n[j])
  }
}

