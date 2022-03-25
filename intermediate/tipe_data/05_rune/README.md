# Rune Golang

dulu ASCII atau kode standar amerika digunakan untuk pertukaaran informasi. di rune kita menggunakan 7bit untuk mewakili 128 karakter, termasuk huruf besar dan kecil, angka, dan berbaagai tanda baca dan karakter kontrol perangkat. karena ini sejumlah besar penduduk dunia tidak dapat menggunakan sistem penulisan mereka sendiri di komputer. jadi untuk mengatasi maasalah ini unicode diciptakan. ini adalah superset ASCII dan berisi semua karakter yang ada dalam sistem penulisan orld termasuk aksen dan tanda diakritik lainnya, kode kontrol seperti tab dan carriage return, dan memberikan masing-masaing nomor standar yagn sebut titik kode unicode, atau dalam bahasa Golang yaitu rune. jenis rune adalah alias dari int32.

contoh rune

```
♄.
```
ini adalah rune dengan nilai heksadesiml

## rune literal

ini mewakili konstanta rune di mna nilai integer mengenali titik kode unicode. dalam bahasa go, literal rune dinyatakan sebaaai satu atau lebih karaakter yang diapit oleh tanda kutip tunggal seperti 'g', '\t' , dll. adi antara tanda kutip tunggal kita diperbolehkan menempatkan karakter apapun kecuali baris baru dan tanda kutip tunggal yang tidak lolos.

```golang
package main

import (
    "fmt"
    "reflect"
)

func main() {
    // membuat rune
    rune1 := 'B'

    // menampilkan rune dan tipe nya
    fmt.Printf("rune 1:  %c; unicode: %U; tipe:%s", rune1, rune1, relflect.TypeOf(rune1))
}
```
