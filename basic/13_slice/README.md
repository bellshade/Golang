# slice

golang slice adalah abstraksi dari golang array, golang array memungkinkan kita untuk mendifinisikan variabel yang dapat menampung beberapa item data dari jenis yang sama tetapi tidak menyediakan metode bawaan apa pun untuk meningkatkan ukurannya secara dinamis atau mendapat sub-arraynya sendiri.

## medifinisikan slice

untuk mendifinisikan sebuah slice, kita dapat mendklarasikan sebagai sebuah array tanpa menentukan ukurannya, atau kita dapat menggunakan fungsi ``make`` untuk membuat slice.

```golang
var angka []int
angka = make([]int, 5, 5) /* memiliki panjang 5 dan kapasitas 5 */
```

selain itu, kita juga dapat membuat slice dari array yang sudah ada sebelumnya tanpa menggunakan function `make()`. dengan cara ini, maka slice yang kita buat akan menjadi *reference* ke array tersebut.

metode ini dilakukan dengan sintaksis `low:high` yang dapat dibagi menjadi 4 jenis:

```golang
array[low:high] // membuat Slice dari array index low hingga index sebelum high (high - 1)
array[low:] // membuat Slice dari array index low hingga index terakhir di array
array[:high] // membuat Slice dari array index 0 hingga index sebelum high (high - 1)
array[:] // membuat Slice dari array index 0 hingga index terakhir di array
```

berikut contohnya:

```golang
// bikin array
days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// bikin slice dari array
weekend := days[5:]   // [low:] dari index ke-5 sampai index terakhir
weekdays := days[:5]  // [:high] dari index ke-0 sampai ke 4 (5 - 1)
codingDays := days[:] // [:] dari index ke-0 sampai index terakhir
meetingDays := days[2:4] // [low:high] index ke-2 sampai index ke 3 (4 - 1)

// print data slice
fmt.Println(weekend)
fmt.Println(weekdays)
fmt.Println(codingDays)
fmt.Println(meetingDays)
```

**output:**

```bash
[Saturday Sunday]
[Monday Tuesday Wednesday Thursday Friday]
[Monday Tuesday Wednesday Thursday Friday Saturday Sunday]
[Wednesday Thursday]
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
