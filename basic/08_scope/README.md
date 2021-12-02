# scope

lingkup dalam pemograman apaun adalah wilayah program di mana variabel yang ditentukan dapat ada dan diluar itu variabel tidak dapat diakses. ada tiga tempat di mana variabel dapat dideklarasikan.

- diluar fungsi atau blok (variabel **lokal**)
- duluar luar semua fungsi (variabel **global**)
- dalam definisi paremeter fungsi (parameter **formal**)

## variabel lokal

variabel yang didekalarasikan di dalam fungsi atau blok disebut variabel lokal. mereka hanya dapat digunakan  oleh pernyataan yang ada di dalm fungsi atau blok kode tersebut.

```golang
package main

import "fmt"

fun main(){
    // contoh dari variabel lokal
    var a, b, c int

    a = 20
    b = 30;
    c = a + b;
    fmt.Printf("totalnya adalah %d", c)
}
```

## variabel global

variabel global didefiinisikan di luar fungsi, biasanya di atas program. variabel global memegang nilainya sepanjang masa program dan mereka dapat diakses di dalam salah satu fungsi yang ditentukan program.

sebuah variabel global dapat diakses oleh fungsi apapun. artinya variabel globak tersedia untuk digunakan di seluruh program setelah deklarasinya. contoh berikut menggunakan variabel global dan lokal.

```golang
package main

import "fmt"

// deklarasi variabel global
var angka int

func main(){
    // variabel lokal
    var a, b int
    
    a = 10
    b = 30
    angka = a + b

    fmt.Printf("angka pertama %d, angka kedua %d", a, b)
    fmt.Printf("total penjumlahan a dan b adalah %d", angka)
}
```

## parameter formal

parameter formal diperlakukan sebagai variabel lokal di dalam fungsi itu dan mereka lebih disukai daripada variabel global misalnya.

```golang
package main

import "fmt"

// membuat variabel global
var a int = 20

func main(){
    var a int = 10
    var b int = 20
    var c int

    c = tambah(a, b)
    fmt.Printf("hasil %d", c)
}

func tambah(a, b int) int{
    return a + b
}
```

