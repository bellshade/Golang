# variabel

Variabel tidak lain adalah nama yang diberikan ke area penyimpanan yang dapat dimanipulasi oleh program. Setiap variabel di Go memiliki tipe spesifik, yang menentukan ukuran dan tata letak memori variabel, rentang nilai yang dapat disimpan dalam memori tersebut, dan rangkaian operasi yang dapat diterapkan ke variabel.

Nama variabel dapat terdiri dari huruf, angka, dan karakter garis bawah. Itu harus dimulai dengan huruf atau garis bawah. Huruf besar dan huruf kecil berbeda karena Go peka terhadap huruf besar-kecil.

## definisi variabel pada go

Definisi variabel memberi tahu kompiler di mana dan berapa banyak penyimpanan yang harus dibuat untuk variabel. Definisi variabel menentukan tipe data dan berisi daftar satu atau lebih variabel dari tipe itu sebagai berikut

```
var nama_variabel tipe_data;
```

Di sini, optional_data_type adalah tipe data Go yang valid termasuk byte, int, float32, complex64, boolean atau objek yang ditentukan pengguna, dll., dan variable_list dapat terdiri dari satu atau lebih nama pengenal yang dipisahkan dengan koma. Beberapa deklarasi yang valid ditampilkan di sini.

```
var angka1, angka2, angka3 int;
var c, x byte;
var angka4, uang float32;
d = 32;
```

Variabel dapat diinisialisasi (diberi nilai awal) dalam deklarasinya. Jenis variabel secara otomatis dinilai oleh kompiler berdasarkan nilai yang diteruskan ke sana. Inisialisasi terdiri dari tanda sama dengan diikuti oleh ekspresi konstan sebagai berikut

```
nama_variabel = value;
```
contoh
```
// deklarasi dari x dan angka2 adalah otomatis int
x = 3, angka2 = 5;
```

## deklarasi tipe statis di go

Deklarasi variabel tipe statis memberikan jaminan kepada kompiler bahwa ada satu variabel yang tersedia dengan tipe dan nama yang diberikan sehingga kompiler dapat melanjutkan kompilasi lebih lanjut tanpa memerlukan detail lengkap variabel. Sebuah deklarasi variabel memiliki arti pada saat kompilasi saja, compiler membutuhkan deklarasi variabel yang sebenarnya pada saat menghubungkan program.

contoh

```golang
package main

import "fmt"

func main(){
  var angka float64
  angka = 42.5
  
  fmt.Println(angka)
  fmt.Printf("tipe variabel angka adalah %T\n", angka)
}
```

## deklarasi tipe dinamis

Deklarasi variabel tipe dinamis membutuhkan kompiler untuk menginterpretasikan tipe variabel berdasarkan nilai yang diteruskan ke sana. Kompiler tidak memerlukan variabel untuk memiliki tipe statis sebagai persyaratan yang diperlukan.

variabel telah dideklarasikan tanpa tipe apa pun. Perhatikan, dalam kasus inferensi tipe, kami menginisialisasi variabel y dengan operator ``:=``, sedangkan x diinisialisasi menggunakan operator ``=``.

```golang
package main

import "fmt"

func main(){
  var angka_saya float64 = 29.9
  angka_lain := 32
  
  fmt.Println(angka_saya)
  fmt.Println(angka_lain)

  fmt.Printf("tipe data angka_saya adalah %T\n", angka_saya)
  fmt.Printf("tipe data angka_lain adalah %T\n", angka_lain)
}
```

## deklarasi variabel campuran go

Variabel dari tipe yang berbeda dapat dideklarasikan sekaligus menggunakan inferensi tipe.

```golang
package main

import "fmt"

func main() {
   var a, b, c = 3, 4, "bellshade"  
	
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a adalah tipe %T\n", a)
   fmt.Printf("b adalah tipe %T\n", b)
   fmt.Printf("c adalah tipe %T\n", c)
}

```
