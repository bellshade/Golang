# bool

tipe data boolean (``bool``) dapat berupa salah satu dari dua nilai, ``benar`` atau ``salah``. boolean digunakan dalam pemograman untuk membuat perbandingan dan untuk mengontrol aliran program.

boolean mewakili nilai kebenaran yang terkait dengan cabang logika matematika, yang menginformasikan algoritma dalam ilmu komputer. dunamakan untuk matematikawan george boole, kata boolean selalu dimulai engan huruf besar ``B``.

tipe data di go for boolean adalah ``bool``. semua huruf kecil. nilai ``true`` dan ``false`` akan selalu dengan huruf kecil ``t`` dan ``f`` masing-masing, karena merupakan nilai khusus di Go.

## operator perbandingan

dalam pemograman, _operator perbandingan_ digunakan untuk membandingkan nilai dan mengevaluasi ke nilai boolean tunggal baik benar atau salah.

| operator      | deskripsi |
| ----------- | ----------- |
| ==      | sama dengan       |
| !=   | tidak sama dengan        |
| <     | kurang dari |
| >     | lebih besar dari |
| <=     | kurang dari atau sama dengan |
| >=     | lebih dari atau sama dengan |

```go
x := 4
y := 7
```
dalam contoh ini, karena nilai ``x`` memiliki nilai ``4``, itu kurang dari ``y`` yang memiliki nilai ``7``

dengan menggunakan kedua variabel tersebut dan nilai terkaitnya, mari kita menulusuri oeprator dari tabel sebelumnya, dalam program ini, kita akan meminta golang untuk mencetak apakah setiap operator perbandingan bernilai benar atau salah.

```go
package main

import "fmt"

func main() {
    x := 4
    y := 7

    fmt.Println("x == y :", x == y)
    fmt.Println("x != y :", x == y)
    fmt.Println("x < y :", x < y)
    fmt.Println("x > y :", x > y)
    fmt.Println("x <= y :", x <= y)
    fmt.Println("x >= y :", x >= y)
}
```
meskipun bilangan bulat digunakan disini. kita dapat menggantinya dengan nilai float. string juga dapat digunakan dengan operator boolean. mereka peka huruf besar/kecil kecuali kita menggunakan metode string tambahan

```go
pace := "pace"
Pace := "Pace"

fmt.println("pace == Pace :" pace == Pace)
```

kita juga dapat menggunakan operator perbandingan lain termasuk ``>`` dan ``<`` untuk membandingkan dua string. Go akan membandingkan string ini secara lekikografis menggunakan nilai karakter ASCII.

kita juga dapat mengevaluasi nilai Boolean dengan operator perbandingan

```go
benar := true
salah := false

fmt.Println("benar != salah :", benar != salah)
```

## operator logika

ada dua operator logika yang digunakan untuk membandingkan nilai. mereka mengevaluasi ekspresi hingga nilai boolean, mengembalikan salah satu ``true`` atau ``false``. operator ini adalah ``&&``, ``||``, dan ``!``

- ``&&`` adalah operator ``and``. benar jika kedua pernyataan ini benar.
- ``||`` adalah operator ``or``. benar jika setidaknya satu pernyataan benar.
- ``!`` adalah operator ``not``. benar hanya jika pernyataan itu salah.

Operator logika biasanya digunakan untuk mengevaluasi apakah dua atau lebih ekspresi benar atau tidak benar. Misalnya, mereka dapat digunakan untuk menentukan apakah nilai lulus dan siswa terdaftar di mata pelajaran, dan jika kedua kasus benar, maka siswa akan diberi nilai dalam sistem. Contoh lain adalah menentukan apakah pengguna adalah pelanggan aktif yang valid dari toko online berdasarkan apakah mereka memiliki kredit toko atau telah melakukan pembelian dalam 6 bulan terakhir.

```go
fmt.Println(( 9 > 7) && ( 2 < 4)) // keduanya benar maka true
fmt.Println((8 == 8) || (6 != 6)) // salah satu pernyataan benar
fmt.Println(!(3 <= 1)) // ekspresi aslinya salah
```

dalam contoh ini

- ``and`` harus memiliki setidaknya satu ekspresi false yang bernilai false
- ``or`` harus memiliki kedua ekspresi yang bernilai false.
- ``!`` harus memiliki ekspresi dalam true agar ekspresi baru bernilai false

kita juga dapat menulis pernyataan majemuk menggunakan ``&&``, ``||`` dan ``!``

```go
!((-0.2 > 1.5) && ((0.8 < 3.1) || (0.1 == 0.1)))
```

lihatlah ekspresi terdalam dahulu ``(0.8 < 3.1) || (0.1 == 0.1)``. ekspresi ini bernilai ``true`` karena kedua pernyataan matematika adalah ``true``.

selanjutanya, Go mengambil nilai yang dikembalikan ``true`` dan menggabungkannya dengan ekspresi dalam berikutnya ``(-0.2 > 1.4) && (true)``. contoh ini kembali ``false`` karena pernyataan matematika ``-0.2 > 1.4`` salah dan ``false`` dan ``true`` kembali ``false``.

akhirnya, kita memiliki ekspresi luar ``!(false)`` yang dievaluasi menjadi ``true``.

## menggunakan operator boolean untuk kontrol aliran

Untuk mengontrol aliran dan hasil program dalam bentuk pernyataan kontrol aliran, kita dapat menggunakan kondisi yang diikuti oleh klausa .

Suatu kondisi mengevaluasi ke nilai Boolean benar atau salah, menyajikan titik di mana keputusan dibuat dalam program. Artinya, suatu kondisi akan memberi tahu kita jika sesuatu bernilai benar atau salah.

Klausa adalah blok kode yang mengikuti kondisi dan menentukan hasil program. Artinya, ini adalah bagian "lakukan ini" dari konstruksi "Jika ``x`` ya true, maka lakukan ini."

```go
if grade >= 65 {
    fmt.Println("lewat")
} else {
    fmt.Println("tidak lewat")
}
```
program ini akan mengevaluasi apakah nilai setiap siswa lulus atau tidak. dalam kasus seorang siswa dengan nilai ``83``, pernyataan pertama akan dievaluasi menjadi ``true``, dan pernyataan cetak ``lewat`` akan dipicu. dalam kasus nilai dibawah 65, pernyataan pertama akan dievaluasi ke ``false``, sehingga program akan melanjutkan untuk mengesekusi pernyataan cetak yang terkait dengan ekpresi else : ``tidak lewat``
