# slice

slice ada referensi elemen array.Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

## inisialisasi slice

Cara pembuatan slice mirip seperti pembuatan array, bedanya tidak perlu mendefinisikan jumlah elemen ketika awal deklarasi.

```go
var makanan = []string{"ayam", "ikan", "udang"}
fmt.Println(makanan[1]) // ikan
```

Salah satu perbedaan slice dan array bisa diketahui pada saat deklarasi variabel-nya, jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah slice.

```go
var makananPertama = []string{"ayam", "bayam"}
var makananKedua = []string{"ikan", "tomat"}
var makananKetiga = []string{"susu", "teh"}
```

## operasi slice

slice dan array tidak bisa dibedakan karena merupakan sebuah kesatuan. Array adalah kumpulan nilai atau elemen, sedang slice adalah referensi tiap elemen tersebut.

Slice bisa dibentuk dari array yang sudah didefinisikan, caranya dengan memanfaatkan teknik 2 index untuk mengambil elemen-nya.

```go
var makanan = []string{"ayam", "ikan", "telur", "daging", "mie"}
var makananSaya = makanan[0:2]

fmt.Println(makananSaya)
```

Kode `makanan[0:2]` maksudnya adalah pengaksesan elemen dalam slice `makann` yang dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2. Elemen yang memenuhi kriteria tersebut akan didapat, untuk kemudian disimpan pada variabel lain sebagai slice baru.
