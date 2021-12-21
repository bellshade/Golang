# Range

kata kunci **range** digunakan dalam **for** loop untuk mengulangi item dari array, slice, channel atau map. dengan array slice, ia mengembalikan indeks item sebagai bilangan bulat.

contoh penggunaan range

```golang
package main

import "fmt"

func main(){
  // membuat slice
  angka := []int{0, 1, 2, 3, 4, 5, 6}

  // print angka
  for i:= range angka{
    fmt.Printlnn("sice item ", i, "adalah", angka[i])
  }
  
  // membuat map
  ibukotaNegara := map[string] string {"indonesia": "jakarta", "jepang": "tokyo"}

  // print ibukotaNegara menggunakan key
  for negara := range ibukotaNegara{
    fmt.Println("ibukota dari negara", negara, "adalah", ibukotaNegara[negara])
  }

  // print menggunakan key value
  for negara, ibukota := range ibukotaNegara{
    fmt.Println("ibukota negara dari", negara, "adalah", ibukota)
  }
}
```
