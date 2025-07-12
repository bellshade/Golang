# Array

Bahasa pemrograman golang menyediakan struktur data yang disebut **array**, yang dapat menyimpan kumpulan elemen berurutan berukuran tetap dari jenis yang sama.

Array digunakan untuk menyimpan kumpulan data, tetapi seringkali lebih berguna untuk menganggap array sebagai kumpulan variabel dengan tipe yang sama.

## mendeklarasikan array

untuk mendeklarasikan array di go, programmer menentukan jenis elemen dan jumlah elemen yang dibutuhkan oleh array.

```
var nama_variable [ukuran] tipe_data
```

```golang
var contoh [10] float32
```

## inisialisasi array

kita dapat menginisialisasi array di go satu per satu atau menggunakan satu pernyataan.

```golang
var contoh = [5] float32 {20.3, 30.2, 55.2, 44.1, 99.2}
```

jumlah nilai di antara kurung kurawal `{}` tidak boleh lebih besar dari jumlah elemen yang kita nyatakan untuk array di antara kurung siku

jika kita menghilangkan ukuran array, array yang cukup besar untuk inisialisasi akan dibuat.

```golang
var contoh [] float32 {20.3, 22.1, 33.5, 72.1, 60.0}
```

## mengakses elemen array

sebuah elemen diakses dengan mengindeks nama array. ini dilakukan dengan menempatkan indeks elemen dalam tanda kurung siku setelah nama array.

```golang
float nilai = contoh[2]
```

pernyataan diatas mengambil elemen ke-3 dari array dan menetapkan nilai ke variabel nilai.

contoh pengunaan secara menyeluruh

```golang
package main

import "fmt"

func main(){
    // membuat array yang terdiri dari
    // 15 angka integer
    var n [15] int
    var i, j int

    // inisialisasi elemen array n ke 0
    for i = 0; i < 15; i++{
        // set elemen pada lokasi i ke i + 3
        n[i] = i + 3
    }
    
    // melihat hasil elemen pada array
    for j = 0; j < 15; j++{
      fmt.Printf("array[%d] = %d\n", j, n[j])
    }
    
  }
```

## slice

Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama. lanjut [disini](SLICE.md)
