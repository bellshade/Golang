# slice

golang slice adalah abstraksi dari golang array, golang array memungkinkan kita untuk mendifinisikan variabel yang dapat menampung beberapa item data dari jenis yang sama tetapi tidak menyediakan metode bawaan apa pun untuk meningkatkan ukurannya secara dinamis atau mendapat sub-arraynya sendiri.

## medifinisikan slice

untuk mendifinisikan sebuah slice, kita dapat mendklarasikan sebagai sebuah array tanpa menentukan ukurannya, atau kita dapat menggunakan fungsi ``make`` untuk membuat slice.

```golang
var angka []int
angka = make([]int, 5, 5) /* memiliki panjang 5 dan kapasitas 5 */
```

## fungsi ken() dan cap()

slice adalah abstraksi dari array. ini sebenarnya menggunakan array sebagai struktur yang mendasarinya. fungsi ``len()`` mengembalikan elemen yang ada dalam irisan di mana fungsi ``cap()`` mengembalikan kapasitas slice (yaitu, beberapa banyak elemen yang dapat di akomodasi). contoh berikut

```golang
package main
import "fmt"

func main(){
  var angka = make([]int, 3, 5)
  /* memanggil fungsi untuk print slice */
  printSlice(angka)

}
func printSlice(x []int){
  fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
```

## slice nihil

jika sebuah slice dideklarasikan tanpa input, maka secara default. ini diinasialisasi sebagai nil. panjang dan kapasitasnya nol.

```golang
package main
import "fmt"

func main(){
  var angka []int
  /* memanggil fungsi untuk print slice */
  printSlice(angka)

  if (angka == nil){
    fmt.Printf("angka nil")
  }
}
func printSlice(x []int){
  fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
```

## sub slicing

slicing memungkinkan batas bawah dan batas atas ditentukn untuk mendapatkan sub slice menggunakan ``[index:index]``

```golang
package main

import "fmt"

func main(){
  /* membuat slice */
  angka := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
  printSlice(angka)

  /* print original slice */
  fmt.Println("angka ==", angka)
  
  /* print dari subslice yng dimulai dari index 1 */
  fmt.Prinfln(angka[1:4] ==, angka[1:4])

  fmt.Println("angka[4:] ==", angka[4:])

  angka1 := make([]int, 0,5)
  printSlice(angka1)

  angka2 := angka[2:]
  PrintSlice(angka2)

  angka3 := angka[2:5]
  printSlice(angka3)
}
func printSlice(x []int){
  fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
```

## fungsi append() dan copy()

kita dapat meningkatkan kapasitas slice menggunakan fungsi ``append()``, menggunakan fungsi ``copy()``, isi dari slice sumber bisa disalink ke tujuan.

```golang
package main

import "fmt"

func main(){
  var angka []int
  printSlice(angka)

  /* append nile slice */
  angka = append(angka, 0)
  printSlice(angka)

  /* menambahkan 1 elemen ke dalam slice */
  angka = append(angka, 1)
  printSlice(angka)

  /* menambahkan lebih dari satu elemen */
  angka = append(angka, 2,3,4)
  printSlice(angka)
  
  /* membuat slice angka1 dengan kapasitas gkita */
  angka1 := make([] int, len(angka), (cap(angka)) * 2)

  /* menyalin konten dari angka ke angka1 */
  copy(angka1, angka)
}

func printSlice(x []int){
  fmt.Printf("len%d cap=%d slice=%v\n", len(x), cap(x), x)
}
```
