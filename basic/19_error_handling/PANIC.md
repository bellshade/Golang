## kapan panic harus digunakan

salah satu faktor penting adalah kita harus menghindri kepanikan dan memulihkan serta menggunakan kesalhan jika memungkinkan. hanya dalam kasus di mana program tidak dapat melanjutkan ekseskusi, mekanisme `panic`dan `recover` harus digunakan.

- kesalahan yng tidak dapat dipulihkan di mana program tidak bisa begitu saja melanjutkan eksekusinya. salah satu contohnya adalah server web yang gagal mengikat ke port yang diperlukan. dalam hal ini wajar untuk panic karena tidak ada lagi yang bisa dilakukan jika pengikatan port itu sendiri gagal.

- katakanlah kita memiliki _metode_ yang menerima pointer sebagai parameter dan seseorang memanggil metode ini menggunakan argumen `nil`. dalam hal ini, `panic` merupakan kesalahan programmer untuk memanggil metode dengan `nil` yang mengharapkan pointer yaang valid

contoh

```go
package main

import "fmt"

func namaLengkap(namaDepan *string, namaBelakang, *string) {
  if namaDepan == nil {
    panic("runtime error: nama depan tidak boleh nil")
  }
  if namaBelakang == nil {
    panic("runtime error: nama belakang tidak boleh nil")
  }

  fmt.Printf("%s %s\n", *namaDepan, *namaBelakang)
  fmt.Println("mengembalikan nilai normal")
}

func main() {
  namaDepan := "majed"
  namaLengkap(&namaDepan, nil)
  fmt.Println("mengembalikan nilai normal")
}
```
