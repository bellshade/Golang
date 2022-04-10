# constant

Constant mengacu pada nilai tetap yang tidak dapat diubah oleh program selama esekusinya. nilai-nilai tetap ini juga disebut dengan **literal**. kostanta dapat berupa salah satu tipe data dasar seperti _konstanta integer_, _kostanta float_, _kostanta karakter_, atau _string literal_. ada juga kostanta enumerasi juga.

Kostanta diperlakukan seperti variabel biasa kecuali bahwa nilainya tidak dapat diubah setelah definisinya.

## literal bilangan bulat

Literal integer dapat berupa kostante desimal, oktal atau heksadesimal. awalan menentukan basis atau radix: ``0x`` atau ``0X`` untuk heksadesimal, 0 untuk oktal, dan tidak ada untuk desimal.

Literal integer juga dapat memiliki sufiks yang merupakan kombinasi dari U dan L, masing-masing untuk ``unsigned`` dan ``long``. akhiran dapat berupa huruf besar atau kecil dan dapat dalam urutan apapun.

Berikut contoh dari literal integer
```
212     // legal
215U    // legal
0xfee   // legal
077     // ilegal: 8 tidak termasuk bilangan oktal
032U    // ilegal: tidak bisa mengulang sufiks
```

Berikut adalah contoh lain dari berbgai jenis literal integer

```
86      // desimal
0123    // oktal
0x4b    // heksadesimal
30      // int
30U     // unsigned int
30l     // long
30ul    // unsigned long
```

## literal float

Literal float memiliki bagian bilangan bulat, titik desimal, bagian pecahan dan bagian eksponen. kita dapat mewakili literal floating point baik dalam bentuk desimal atau bentuk eksponensial.

Saat mewakili menggunakan bentuk desimal kita harus menyertakan titik desimal, eksponen, atau keduanya dan saat mewakili menggunakan bentuk eskponensial, anda harus menyertakan bagian bilangan bulat, bagian pecahan, atau keduanya. Eksponen bertanda diperkenalkan oleh ``e`` atau ``E``.

Berikut beberapa contoh literal float
```
3.142152       // legal
5125225-5L     // legal
510E           // ilegal: exponent tidak ada
120f           // ilegal tidak desinak eksponensial
.e55           // ilegal tidak ada integer setelah titik
```

| escape string | arti |
| :---: | :--: |
| \\ | karakater \ | 
| \' | karakter '  |
| \" | karakter "  |
| \A | alert atau peringatan |
| \B | mengapus     |
| \F | form         |
| \n | baris baru   |
| \R | carriage     |
| \t | indent horizontal |
| \v | vertival indent |
| \ooo | bilangan okal satu sampai tiga digit |
| \xhh... | bilangan heksadesimal dari satu atau lebih digit |

Contoh pengunaan ``\t`` pada pogram

```golang
package main

import fmt

func main(){
    fmt.Printf("bell\thade")
}
```

## keyword const

Kita dapat menggunakan awalan cost untuk mendeklarasikan kostanta dengan tipe tertentu dengan contoh:

```
const nama_variabel tipe_data = value;
```

Contoh berikut penggunaan const

````golang
package main

import fmt

func main(){
    const PANJANG = 12
    const LEBAR = 2

    var hasil int
    hasil = PANJANG * LEBAR;

    fmt.Printf("hasil: %d", hasil)
}
```