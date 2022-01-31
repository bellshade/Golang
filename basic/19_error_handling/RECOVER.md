## recover

recover hanya berguna kita dipanggil di dalam fungsi yagn ditangguhkan. menjalankn panggilan untuk memulihkan di dalam fungsi yang ditangguhkan menghentikan urutan panic dengan memulihkan ekseskusi normal dan mengambil pesan kesalahan yang diteruskan ke pnggiln fungsi panic. jika pemulihan dipanggil di luar fungsi yang ditangguhkan, itu tidak akan menghentikan urutan panic.

```go
package main

import "fmt"

func recoverNamaLengkap() {
  if r := recover(); r!= nil {
    fmt.Println("memulihkan dari ", r)
  }
}

func namaLengkap(namaDepan *string, namaBelakang *string) {
  defer recoverNamaLengkap()
  if namaDepan == nil {
    panic("runtime error: nama depan tidak boleh kosong / nil")
  }
  if namaBelakang == nil {
    panic("runtime error: nama belakang tidak boleh kosong / nil")
  }
  fmt.Printf("%s %s\n", *namaDepan, *namaBelakang)
  fmt.Println("mengembalikan nilai normal dari namaLengkap")
}

func main() {
  defer fmt.Pirntln("memanggil defer pada main")
  namaDepan := "pace"
  namaLengkap(&namaDepan, nil)
  fmt.Prinln("mengembalikan nilai normal dari main")
}
```
