# Error handling

pemograman go menyediakan kerangka kerja penanganan kesalahan yang cukup sederhan dengan tipe interface kesalahan bawaan dari deklarasi berikut:

```
type error interface {
  Error() string
}
```

fungsi biasany mengembalikan kesalahan sebagai nili pengembalian terakhir. gunakan ``errors.New`` untuk membuat pesan kesalahan dasar sebgai berikut

```go
func Sqrt(value float64) (float64, error) {
  if (value < 0) {
    return 0, errors.New("Math: angka tidak boleh kurang dari 0")
  }
  return math.Sqrt(value), nil
}
```

gunakan nilai kembali dan pesan kesalahan

```go
hasil, err := Sqrt(-1)

if err != nil {
  fmt.Println(err)
}
```

contoh

```go
package main

import "errors"
import "fmt"
import "math"

func Sqrt(value float64) (float64, error) {
  if(value < 0) {
    return 0, errors.New("math: angka negative")
  }
  return math.Sqrt(value), nil
}

func main() {
  result, err := Sqrt(-1)
  
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(result)
  }

  result, err = Sqrt(9)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(result)
  }
}
```

## defer 

dalam bahasa go, pernyataan defer menunda ekseskusi fungsi atau metode anonim hingga fungsi terdekat kembali. dengan kata lain, argumen penangguhan fungsi atau pemanggilan metode dievaluasi secara instan, tetapi mereka tidak mengekseskusi saampai funsi terdekat kembali. kita dapat membuat metode yang ditangguhkan, atua fungsi, atau fungsi anonim dengan menggunakan keyword ``defer``

contoh lebih lanjut dari [defer](DEFER.md)

## panic

cara idiomatis untuk menangani kondisi abnormal dalm progrma go adalah menggunakan error. kesalahan cukup untuk sebagian besr kondisi abnormal timbul dalam program.

tetapi beberap situasi di mana program tidak dapat melanjutkan ekseskusi setelah kondisi abnormal. dalam hasil ini, kita menggunakan ``panic`` untuk menghentikan program sebelum waktunya. ketika suatu fungsi mengalmi kepanikan, eksekusiny dihentikan, semua fungsi yang ditangguhkn dijalankan dan kemudian kontrol kembali ke pemanggilannya. proses ini berlanjut sampai emua fungsi goroutine saat ini telh kembali di mana program mencetak pesan panic, diikuti oleh pelacakan  tumpukan dan kemudian dihentikan

lebih lanjut dari [panic](PANIC.md)

## recover

_recover_ adalah fungsi bawaan yang digunakan untuk mendapatkan kembali kendali atas program yang panic.

```
func recover() interface()
```

lebih lanjut dari [recover](RECOVER.md)

