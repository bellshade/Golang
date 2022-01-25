# Type casting

konversi tipe adalah cara untuk megubah suatu variabel dari satu tipe data ke tipe data lainnya. Misalnya, jika anda ingin menyimpan panjang ke dalam bilangan int sederhana, anda bisa mengetikkan _type cast long to int_. anda dapat mengkonversi nilai dari satu jenis ke jenis lainnya menggunakan operator cast.

```
tipe(ekspresi)
```

## contoh

contoh berikut di mana operator cast menyebabkan pembagian satu variabel integer dengan uang lain dilakukan sebagi operasi bilangan float.

```golang
package main

import "fmt"

func main() {
  var hasil int = 17
  var hitung int = 5
  var ubah float32

  ubah = float32(hasil) / float32(hitung)
  
  fmt.Printf("hasil dari valye ubah adalah: %f\n", ubah)
}
```

golang juga tidak meperbolehkan menggunakan tipe data yang berbeda dalam melakukan operasi.contoh

```golang
package main

import "fmt"

func main() {
  var angka int = 19
  var angkaLain float32 = 21
  var hasil float32

  hasil = angka * angkaLain
  
  fmt.Printf("hasilnya adalah %f\n", hasil)
}
```

pada saat dikompilasi maka akan terjadi error dikarenakan kedua tipe data dari variabel yang berbeda
