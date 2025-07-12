## defer

contoh sintaks dari defer

```go
//fungsi

defer func nama_fungsi(list_parameter Type) return_tipe {
  // kode
}

// method
defer func (tipe penerima) nama_method(lis_parameter) {
  // kode
}

defer func (list_parameter) (tipe_return) {
  // kode
}()
```

- dalam bahasa go, beberapa pernyatan defer diperbolehkan dalam program yang ama dan mereka dieksekusi dalam urutan LIFO (_last in first out_)

- dalam pernyataan defer, argumen dievaluasi saat pernyataan penangguhan dieksekusi, bukan saat dipanggil.

- pernyataan penangguhan umumny digunakan untik memastukan bhwa file ditutup ketika kebutuhan mereka selesai, atau untuk menutup saluram, atau menangkap panic dalam program

contoh dri fungsi defer

```go
package main

import "fmt"

func perkalian(angka1, angka2 int) int {
  hasil := angka1 * angka2
  fmt.Println("hasil :", hasil)

  return 0
}

func tampilKata() {
  fmt.Println("bellshade")
}

func main() {
  // memanggil fungsi perkalian
  perkalian(25, 4)

  // memanggil fungsi perkalian
  // dengan menggunakan defer
  defer perkalian(7, 4)

  tampilKata()
}
```

- pertama kita memanggil fungsi `perkalian` secara normal (tanpa kata kunci `defer`), ketika dijalnkn fungsi dipanggil dengan hasil 25 * 4 = 100.

- kedua kita memanggil fungsi `perkalian` sebagai fungsi defer menggunakan kata kunci `defer` dan dijalankan dengan hasil 7 * 4 = 28, ketika semua metode di sekitarnya kembali
