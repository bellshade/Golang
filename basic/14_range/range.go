package main

import "fmt"

func main(){
  // membuat slice
  angka := []int{0, 1, 2, 3, 4, 5}
  
  // print angka
  for i := range angka{
    fmt.Println("slice item", i, "adalah", angka[i])
  }

  // membuat map
  ibukotaNegara := map[string] string{
    "indonesia": "jakarta",
    "italia": "roma",
    "jepang": "tokyo",
    "russia": "moskow",
  }
  
  fmt.Println("\nMenggunakan key")
  // print menggunakan key
  for negara := range ibukotaNegara{
    fmt.Println("ibukota negara dari", negara, "adalah", ibukotaNegara[negara])
  }

  fmt.Println("\nMenggunakan key valu")
  // print menggunakan key value
  for negara, ibukota := range ibukotaNegara{
    fmt.Println(
      "ibukota dari negara", negara, "adalah", ibukota)
  }
}

