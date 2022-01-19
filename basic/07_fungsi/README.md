# Fungsi 

fungsi adakah sekelompok pernyataan yang bersama-sama melakukan tugas. setiap program golang memiliki setidaknya memiliki satu fungsi yaitu ``main()``. kitadapat membagi kode kita menjadi fungsi-fungsi terpisah. bagaimana kita dapat membagi kode kita di antara fungsi yang berbeda terserah anda, tetapi secara logis, pembagian harus sedemikian rupa sehingga setiap fungsi melakukan tugas-tugas tertentu

## mendifinisikan fungsi

```
func nama_fungsi( list_parameter ]) [tipe_pengembalian]
{
    isi_dari_fungsi
}
```

- func

    deklarasi suatu fungsi

- nama_fungsi

    ini adalah nama sebenarnya dari fungsi tersebut

- parameter

    parameter seperti placeholder. saat suatu fungsi dipanggil, ita memberikan nilai ke paramter. nilai ini disebut sebagai paramter atau argumen aktual. daftar parameter mengacu pada jenis, urutan, dah jumlah paramter suatu fungsi

- tipe pengembalian

    sebuah fungsi dapat mengembalikan daftar nilai. ``return_types`` adalah daftar tipe data dari nilai yang dikembalikan oleh fungsi. beberapa fungsi melakukan operasi yang diingankan tanpa mengembalikan nilai, dalam hal ini ``return_type`` tidak diperlukan

- badan fungsi

    ini berisi dari kumpulan pernyatan yang mendifinisikan apa yng dilakukan dari fungsi

contoh

```golang
// menghitung kedua variabel
func hitung(angka1, angka2 int) int{
    result int

    result = angka1 + angka2

    return result
}
```

## memanggil fungsi

Ketika sebuah program memanggil fungsi, kontrol program ditransfer ke fungsi yang dipanggil. Fungsi yang dipanggil melakukan tugas yang ditentukan dan ketika pernyataan pengembaliannya dieksekusi atau ketika kurung kurawal penutup fungsi tercapai, ia mengembalikan kontrol program kembali ke program utama.

Untuk memanggil suatu fungsi, anda hanya perlu meneruskan parameter yang diperlukan bersama dengan nama fungsinya. Jika fungsi mengembalikan nilai, maka Anda dapat menyimpan nilai yang dikembalikan.

contoh
```golang
package main

import "fmt"

func main(){
    var a int = 10
    var b int = 20
    var hasil int
    
    hasil = hitung(a, b)

    fmt.Print("hasilnya adalah %d\n", hasil)
    
}

func hitung(angka1, angka2, int) int{
    var hasil int 
    result = angka1 + angka2

    return result
}
```

## Fungsi dengan multiple return

Di bahasa pemprograman golang kita dapat membuat fungsi dengan nilai balik lebih dari satu.cara membuatnya hampir sama dengan funsgi
biasa.Perbedaanya kita hanya perlu perlu menambahkan tipe data untuk return value yang kedua dan seterusnya.Untuk lebih jelasnya 
berikut adalah contoh codenya.

```go
func hitungPersegi(sisi int) (int, int) {
   luas := sisi * sisi
   keliling := sisi * 4
   return keliling, luas
}
```

Untuk memanggil funginya kita harus membuat variabel dengan jumlah yang sama dengan jumlah return fungsinya.

```go
func main(){
   keliling, luas := hitungPersegi()
   fmt.Println("keliling persegi : ", keliling)
   fmt.Println("luas persegi : ", luas)
}
```

Di dunia go, kita nanti akan sering sekali menjumpai fungsi dengan multiple return, baik fungsi-fungsi bawaan dari go atau funsgi 
dari library yang kita pakai.
