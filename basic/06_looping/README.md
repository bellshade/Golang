# Looping

mungkin ada siatuasi, ketika anda perlu mengeksekusi blok kode beberapa kali. secara umum, pernyataan dieksekusi secara berurutan;  pernyataan pertama dalam suatu fungsi dekseskusi terlebih dahulu, diikuit oleh yang kedua dan seterusnya

bahasa pemograman menyediakan berbagai struktur kontrol yang memungkinkan jalur eksekusi yang lebih rumit

pernyataan loop memungkinkan kita untuk mengseskusi pernyataan atau sekelompok pernyataan beberapa kali dan berikut ini adalh contoh dari looping

## for looping

**for looping** ini mengeksekusi urutan pernyataan beberapa kali dan menyingkat kode yang mengelola variabel loop.

```
for [kondisi | (init; kondisi; increment) | range]{
    statement;
}
```

contoh
```golang
package main

import "fmt"
func main(){
    var b int = 15
    var a int
    angka := [6]int{1, 2, 3, 5}

    // esekusi dari for loop
    for a := 0; a<10; a++{
        fmt.Printf("Value dari a: %d\n", a)
    }
    
    for a < b{
        a++
        fmt.Printf("Value dari a: %d\n", a)
    }
    
    for i, x:= range angka{
        fmt.Printf("value dari x = %d di %d\n", x, i)
    }
}
```
alur kendali dalam **for** loop adalah sebagai berikut:
- jika suatu kondisi tersedia, maka for loop dijalankan selama kondisinya benar.
- jika klausa **for** yaitu ``(init; kondisi; increment;)`` ada maka

    - langkah ini dijalankan terlebih dahulu, dan hanya sekali. langkah ini memungkinkan kita untuk mendklarasikan dan menginisialisasi variabel kontrol loop apa pun. kita tidak perlu membuat pernyataan disini, selama titik koma muncul.

    - selanjutnya, kondisinya dievaluasi, jiika benar, maka badan perulangan akan dieksekusi. jika salah, badam perulangan tidak dieksekusi dan aliran kontrol melompat ke pernyataan berikutnya setelah perulangan **for**.

    - setelah tubuh for loop dieksekusi, aliran kontrol melompat kembali ke pernyataan **increment**. pernyataan ini memungkinkan kita untuk memperbarui variabel kontrol loop apa pun. pernyataan ini dapat dikosongkan, selamat titik koma muncul setelah kondisi

    - kondisi tersebut kini deavluasi kembali. jika benar, loop diekseskusi dan proses berulang (body of loop, kemudian langkah increment, dan kemudian kondisi lagi). setelah kondisi menjadi salah, perulangan for berakhir.

## nested looping ( loop bersarang )

bahasa golang memungkinkan untuk menggunakan satu loop dalam loop lain. bagian berikut menunjukkan beberapa contoh untuk mengilustrasikan konsep

```
for [kondisi | (init; kondisi; increment) | range]{
    for [kondisi | (init; kondisi;  increment) | range]{
        statement(s);
    }
}
```
program dibawah ini menggunakan perulangan bersarang untuk mencari bilangan prima dari 2 hingga 100

```golang
package main

import "fmt"

func main(){
    var i, j int
    
    for i = 2; i < 100; i++{
        for j = 2; k <= (i / j); j++{
            if (i%j == 0){
                break;
            }
        }
        if (j > (i / j)){
            fmt.Printf("%d adalah bilangan prima\n", i);
        }
    }
}
```

## loop tak terbatas

sebuah loop menjadi tak terbatas jika kondisinya tidak pernah salah. for loop secara tradisional digunakan untuk tujuan ini. karena tidak ada satu pun dari tiga ekspresi yang membentuk for loop yang diperlukan, kita dapat membuat loop tak berujung dengan membiarkan ekspresi kondisional kosong atau dengan memberikan **true** padanya

```golang
package main

import fmt

func main(){
    for ture{
        fmt.Printf("looping ini akan berjalan selamanya\n")
    }
}
```
