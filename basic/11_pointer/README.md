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

