## string

string adalah urutan satu atau lebih karakter (huruf, angka, simbol) yang dapat berupa konstanta atau variabel. String ada di dalam tanda kutip belakang ` atay tanda kutip ganda " di Go dan memiliki karateristik yang berbeda tergantung pada tanda kutip yang kita gunakan.

jika menggunakan tanda kutip kembali, kita membuat literal string mentah. jika kita menggunakan tanda kutip ganda, kita membuat literal string yang ditafsirkan.

## literal string mentah (raw)

literal string mentah adalah urutan karakter antara tanda kutip kembali, sering disebut kutu kembali. Di dalam tanda kutip, karakter apa pun akan muncul seperti yang ditampilkan di antara tanda kutip belakang, kecuali untuk karakter tanda kutip belakang itu sendiri.

```golang
a := `katakan "tidak" pada kuman`
fmt.Println(a)
```

Biasanya, garis miring terbalik digunakan untuk mewakili karakter khusus dalam string. Misalnya, dalam string yang ditafsirkan, ``\n``akan mewakili baris baru dalam sebuah string. Namun, garis miring terbalik tidak memiliki arti khusus di dalam literal string mentah:

```golang
a := `katakan "iya" pada vitamin`
```

## string dengan karakter UTF-8

``UTF-8`` adalah skema pengkodean yang digunakan untuk mengkodekan karakter lebar variabel menjadi satu hingga empat byte. Go mendukung karakter ``UTF-8`` di luar kotak, tanpa pengaturan, perpustakaan, atau paket khusus apa pun. Karakter Romawi seperti huruf Adapat direpresentasikan dengan nilai ASCII seperti angka 65. Namun, dengan karakter khusus seperti karakter internasional 世, ``UTF-8`` akan diperlukan. Go menggunakan runetipe alias untuk data ``UTF-8``.

```golang
a := "Hello, 世界"
```

kita dapat menggunakan rangekata kunci dalam satu forlingkaran untuk mengindeks string apa pun di Go,untuk saat ini, penting untuk diketahui bahwa kita dapat menggunakan ini untuk menghitung byte dalam string yang diberikan:

```golang
package main

import "fmt"

func main() {
    a := "Hello, 世界"
    for i, c := range a {
        fmt.Printf("%d : %s\n", i, string(c))
    }
    fmt.Println("panjang dari Hello, 世界", len(a))
}
```
