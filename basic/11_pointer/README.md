# pointer

pointer di go mudah dan menyenangkan untuk dipelajari. beberapa tugas pemograman go dilakukan lebih mudah dengan pointer dan tugas lainnya, seperti panggilan dengan referensi, tidak dapat dilakukan tanpa menggunakan pointer. jadi, penting untuk mempelajari untuk menjadi pemogrammer golang yang sempurn.

seperti yang kita ketahui, setiap variable adalah lokasi memori memiliki alamat yang ditentukan yang dapat diakses menggunakan operator ampersand ``&``. yang menunjukkan alamat di memori.

```golang
package main

import "fmt"

func main(){
  var angka int = 20
  fmt.Printf("alamat dari variabel angka adalah: %x\n", &a)
}
```

## apa itu pointer

pointer adalah variabel yang nilainya adalah alamat dari variabel lain, yaitu, alamat langsung dari lokasi memori. Seperti variabel atau konstanta apa pun, kita harus mendeklarasikan pointer sebelum kita dapat menggunakannya untuk menyimpan alamat variabel apa pun.

```
var nama_variable *tipe_data
```

## cara menggunakan pointer

ada beberapa operasi penting, yang sering kita lakukan dengan pointer:

- kita mendifinisikan variabel pointer
- menetapkan alamat variabel ke pointer
- mengakses nilai pada alamat yang disimpan dalam variabel pointer

semua operasi ini dilakukan dengan menggunakan operator unary ``*`` yang mengembalikan nilai variabel yang terletak di alamat yang ditentukan oleh operandnya

```golang
package main

import "fmt"

func main(){
  var angka int = 30
  var poin *int
  
  // menyimpan ke dalam variable pointer
  poin = &angka

  fmt.Printf("alamat dari variabel angka adalah: %x\n", &a)

  // alamat yang kesimpn dalam variabel pointer
  fmt.Printf("alamat yang kesimpan pada variabel pointer: %x\n", poin)
  
  // mengakses value menggunakan pointer
  fmt.Printf("value dari variabel poin adalah: %d\n", *poin)
}
```

## pointer nihil pada golang

Go compiler menetapkan nilai Nil ke variabel pointer jika kita tidak memiliki alamat yang tepat untuk ditugaskan. Ini dilakukan pada saat deklarasi variabel. Pointer yang diberi nil disebut pointer nil .

```golang
package main

import "fmt"

func main(){
  var ptr *int
  fmt.Printf("valua dari ptr adalah: %x\n", ptr)
}
```
Pada sebagian besar sistem operasi, program tidak diizinkan untuk mengakses memori pada alamat 0 karena memori tersebut dicadangkan oleh sistem operasi. Namun, alamat memori 0 memiliki arti khusus; itu menkitakan bahwa penunjuk tidak dimaksudkan untuk menunjuk ke lokasi memori yang dapat diakses. Tetapi menurut konvensi, jika sebuah pointer berisi nilai nil (nol), diasumsikan tidak menunjuk ke apa pun.

untuk memeriksa pointer niil, kita dapat menggunakan pernyataan if sebagai berikut
```golang
// mengecek jika pointer tidak nil
if(variabel_pointer != nil)

// mengecek jika pointer nil
if(variabel_pointer == nil)
```

## Pass By Reference & Pointer Receiver

```
Catatan : Di dalam Golang sebenarnya tidak ada istilah pass by reference namun kita menggunakan istilah pass by reference di sini walau sebenarnya istilah tersebut kurang tepat. Sebab, reference sendiri merujuk pada di mana data suatu vairabel di simpan di dalam memori. Bahasa seperti C++ menggunakan istilah reference untuk variabel-variabel yang memiliki alamat memori yang sama satu dengan yang lain sehingga perubahan pada satu variabel bisa berdampak pada variabel yang lain. Sementara itu, variabel di Golang disimpan dalam alamat memori yang berbeda (unik) sehingga ketika satu variabel mengalami perubahan, perubahan lain tidak akan terpengaruh. Selengkapnya silakan baca di https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
```

Ketika kita membuat function di `Go` di mana kita ingin ketika ada perubahan pada nilai parameter di function tersebut, maka argumen asli dari parameter tersebut juga ikut berubah. Hal ini disebut sebagai *call by reference*, di mana kita menggunakan alamat memori atau pointer dari parameter yang kita lempar ke function. Perhatikan contoh di bawah ini (kode diolah dari [https://www.tutorialspoint.com/go/go_function_call_by_reference.htm](https://www.tutorialspoint.com/go/go_function_call_by_reference.htm))

```Go
func passByReference(p *int, q *int){
  var temp int
  temp = *q
  *q = *p
  *p = temp
}

func main(){
  p := 2
  q := 3

  fmt.Printf("value of p before : %v \n", p)
  fmt.Printf("value of q before : %v \n", q)

  callByReference(&p, &q)

  fmt.Printf("value of p after : %v \n", p)
  fmt.Printf("value of q after : %v \n", q)
}
```

Kita akan mendapatkan hasil seperti ini.

```
value of p before : 2
value of q before : 3
value of p after : 3
value of q after : 2
```

Seperti yang kita lihat, ketika kita melakukan *pass by reference*, dalam hal ini pada function `passByReference`, kita melakukan *pass* terhadap alamat memori atau *pointer* dari variabel `p` dan `q`. Di dalam function tersebut, kita melakukan pertukaran terhadap dua variabel tersebut menggunakan *pointer* dari masing-masing parameter. Hasil dari function `main` adalah kedua nilai variabel ikut berubah. Sehingga, **ketika kita melakukan pass by reference, perubahan yang terjadi di fungsi pemanggil juga terjadi pada fungsi aslinya** (pada kasus ini, fungsi pemanggil adalah `callByReference`, dan fungsi asli adalah `main`)

Sama seperti *call by reference*, *pointer receiver* kita gunakan ketika kita menghendaki adanya perubahan pada method struct juga berlaku pada pemanggilnya. Perhatikan kode berikut (kode diambil dari [https://go.dev/tour/methods/4](https://go.dev/tour/methods/4))

```Go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

Jika kita jalankan, kita akan mendapatkan hasil `50`. Sementara itu, jika method `Scale` kita ubah menjadi `func (v Vertex) Scale(f float64)`, kita akan mendapatkan hasil `5`. Hal ini disebabkan karena ketika kita melakukan *value receiver* (method `Scale` menggunakan receiver `v Vertex`) perubahan terhadap field `X` dan `Y` di dalam method `Scale` tidak tersimpan. Sebaliknya, jika kita melakukan *pointer receiver* (method `Scale` menggunakan receiver `v *Vertex`), yang terjadi adalah perubahan terhadap field `X` dan `Y` di dalam method `Scale` akan tersimpan.

