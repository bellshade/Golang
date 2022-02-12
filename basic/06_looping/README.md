# Looping

Pada situasi tertentu anda akan perlu mengeksekusi sebuah blok kode beberapa kali dan untuk mencapai hal tersebut maka diperlukan sebuah perulangan/looping. Secara umum, pernyataan dieksekusi secara berurutan; pernyataan pertama dalam suatu fungsi dieksekusi terlebih dahulu, diikuti oleh yang kedua dan seterusnya

bahasa pemograman menyediakan berbagai struktur kontrol yang memungkinkan jalur eksekusi yang lebih rumit

pernyataan loop memungkinkan kita untuk mengeksekusi pernyataan atau sekelompok pernyataan beberapa kali dan berikut ini adalah contoh dari penerapan looping pada Go

## for looping

**for looping** ini mengeksekusi urutan pernyataan beberapa kali dan menyingkat kode yang mengelola variabel loop.

```
for [kondisi | (init; kondisi; increment) | range]{
    // code here! (badan perulangan)
}
```

Contoh

```golang
package main

import "fmt"

func main() {
    var b int = 15
    var a int

    // deklarasi array (pembahasan lebih lanjut pada materi 10_array)
    angka := [6]int{1, 2, 3, 5}

    // Cara Pertama
    /*
    1. untuk melakukan perulangan pada Go harus di awali dengan keyword for
    2. deklarasi variabel baru (variabel a pada looping pertama tidak sama dengan variabel a diatas)
    3. a<10, merupakan kondisi logika yang akan dicek setiap kali perulangan
    4. a++ (shorthand dari a + 1), merupakan penambahan nilai ke variabel a di for-loop. Jika setiap kali perulangan menginginkan penambahan dua atau lebih maka digunakan (a + 2), dst. Sebaliknya konsep yang sama pada pengurangan (a--) atau (a - 1).
    5. statement, merupakan blok kode yang akan dijalankan.
    */
    for a := 0; a < 10; a++ {
        fmt.Printf("Value dari a: %d\n", a)
    }

    // Cara Kedua
    // setelah perulangan diatas selesai maka akan dilanjutkan dengan perulangan dibawah
    // variabel yang digunakan pada kondisi adalah dua variabel awal diatas
    // dan penambahan nilai variabel terjadi pada badan perulangan (merupakan alternatif dari cara pertama)
    // cara ini merupakan while loop pada bahasa pemrograman lain
    for a < b {
        a++
        fmt.Printf("Value dari a: %d\n", a)
    }

    // Cara Ketiga
    // penggunaan keyword break dan continue
    // continue akan meng-skip sebuah perulangan
    // sedangkan break akan menghentikan perulahan bahkan sebelum kondisi sampai false
    for i := 10; i > 0; i-- {
        if i%2 == 0 {
            fmt.Printf("skip value dari i = %d\n", i)
            continue
        }

        fmt.Printf("value dari i = %d\n", i)

        if i == 5 {
            fmt.Printf("value dari i = %d, perulangan ketiga berhenti!\n", i)
            break
        }
    }

    // Cara Keempat
    // for - range biasa digunakan untuk melakukan perulangan pada data bertipe array
    // range mengembalikan (return) dua nilai (value) yaitu index (i) dan data dari index tersebut (x)
    for i, x := range angka {
        fmt.Printf("value dari x = %d di %d\n", x, i)
    }
}
```

alur kendali dalam **for** loop adalah sebagai berikut:

- jika suatu kondisi tersedia, maka for loop dijalankan selama kondisinya benar.
- jika klausa **for** yaitu `(init; kondisi; increment;)` ada maka

  - langkah ini dijalankan terlebih dahulu, dan hanya sekali. langkah ini memungkinkan kita untuk mendeklarasikan dan menginisialisasi variabel kontrol loop apa pun. kita tidak perlu membuat pernyataan disini, selama titik koma muncul.

  - selanjutnya, kondisinya dievaluasi, jika benar, maka badan perulangan akan dieksekusi. jika salah, badan perulangan tidak dieksekusi dan alur kontrol melompat ke pernyataan berikutnya setelah perulangan **for**.

  - setelah badan for loop dieksekusi, alur kontrol melompat kembali ke pernyataan **increment**. pernyataan ini memungkinkan kita untuk memperbarui variabel kontrol loop apa pun. pernyataan ini dapat dikosongkan, selamat titik koma muncul setelah kondisi

  - kondisi tersebut kini dievaluasi kembali. jika benar, loop dieksekusi dan proses berulang (body of loop, kemudian langkah increment, dan kemudian kondisi lagi). setelah kondisi menjadi salah, perulangan for berakhir.

## nested looping ( loop bersarang )

bahasa golang memungkinkan untuk menggunakan satu loop dalam loop lain. bagian berikut menunjukkan beberapa contoh untuk mengilustrasikan konsep tersebut

```
for [kondisi | (init; kondisi; increment) | range]{
    // code here!
    for [kondisi | (init; kondisi;  increment) | range]{
        // code here!
    }
}
```

program dibawah ini menggunakan perulangan bersarang untuk mencari bilangan prima dari 2 hingga 100

```golang
package main

import "fmt"

func main() {
    var i, j int

    for i = 2; i < 100; i++ {
        for j = 2; j <= (i / j); j++ {
            if (i%j == 0) {
                break;
            }
        }
        if j > (i / j) {
            fmt.Printf("%d adalah bilangan prima\n", i);
        }
    }
}
```

## loop tak terbatas

sebuah loop menjadi tak terbatas jika kondisinya tidak pernah salah. for loop secara tradisional digunakan untuk tujuan ini. karena tidak ada satu pun dari tiga ekspresi yang membentuk for loop yang diperlukan, kita dapat membuat loop tak berujung dengan membiarkan ekspresi kondisional kosong atau dengan memberikan **true** padanya

```golang
package main

import "fmt"

func main() {
    for true {
        fmt.Printf("looping ini akan berjalan selamanya\n")
    }

    // atau
    /*
    for {
        fmt.Printf("looping ini akan berjalan selamanya\n")
    }
    */
}
```

## recursion

recursion merupakan salah satu konsep pada perulangan yang memanggil fungsi itu sendiri.

```golang
package main

import "fmt"

func main() {
    var n int = 3
    fmt.Printf("%d! = %d", n, factorial_number(n))
}

func factorial_number(n int) int {
    fmt.Printf("bilangan faktorial %d\n", n)

    if n == 0 {
        return 1
    }

    return n * factorial_number(n-1)
}
```

berikut ilustrasinya
```
factorial_number(3)
| 3 - 1
|---factorial_number(2)
|    | 2 - 1
|    |---factorial_number(1)
|    |   | 1 - 1
|    |   |---factorial_number(0)
|    |   |    | n == 0 --> return 1;
|    |   |    factorial_number(1) --> 1 --> 1 * factorial_number(0) = 1 * 1 = 1
|    |   factorial_number(2) --> 2          2 * factorial_number(1) = 2 * 1 = 2
|    factorial_number(3) --> 3              3 * factorial_number(2) = 3 * 2 = 6
```
