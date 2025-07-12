# structure

array golang memungkinkan anada untuk mendefinisikan variabel yang dapat menampung beberapa item data dari jenis yang sama. **structure atau struktur** adalah tipe data lain yang ditentukan pengguna yang tersedia dalam pemograman golang, yang memungkinkan kita untuk menggabungkan item data dari berbagai jenis.

struktur digunakan untuk mempresentasikan record. misalnya kita ingin melacak buku-buku di perpustakaan. kita ingin melacak atribut berikut pada setiap buku

- judul
- pengarang
- subjek
- id buku

## mendifinisikan struktur

untuk mendifinisikan struktur, kita harus menggunakan pernyataan tipe dan **struct** . Pernyataan struct mendifinisikan tipe data baru, dengan banyak anggota untuk program kita. pernyataan tipe mengikat nama dengan tipe yang struct dalam kasus kita.

```golang
type tipe_struct_variabel strcut{
    definisi member;
    deifnisi member;
}
```

atau kita juga bisa menyingkatnya dengan cara

```golang
nama_variabel := tipe_struktur {value1, value2, value3}
```

## mengakses anggota struct

untuk mengakses setiap enggota struktur, kita menggunakan operator akses anggota (.). operator akses anggota dikodekan sebagai periode antara nama variabel struktur dan anggota struktur yang ingin kita akses.

```golang
package main

import "fmt"

type Buku struct{
    judul string
    penulis string
    id_buku int
}

func main(){
    var Buku1 Buku

    Buku1.judul = "Harry potter the dathly hallows"
    Buku1.penulis = "J.K Rowling"
    Buku1.id_buku = 1761
    
    fmt.Printf("judul buku %s\n", Buku1.judul)
    fmt.Printf("dengan penulis %s\n", Buku1.penulis)
    fmt.Printf("id buku adalah %d\n", Buku1.id_buku)
}
```

## sturct sebagai argumen fungsi

kita dapat meneruskan struktur sebagai argumen fungsi dengan cara yang sangat mirip saat kita menersukan variabelatau petunjuk lainnya. kita akan mengakses variabel struktur enganc ara yang sama seperti yang kita lakukan pada contoh di atas

```golang
package main

import "main"

type Buku struct{
    judul string
    penulis string
    id_buku int
}

func main(){
    var Buku1 Buku
    var Buku2 Buku

    Buku1.judul = "harry potter - the goblet of fire"
    Buku1.penulis = "jk rowling"
    Buku1.id_buku = 15241

    Buku2,judul = "harry potter - the prisoner of azkaban"
    Buku2.penuls = "jk rowling"
    Buku2.id_buku = 15242

    tampilBuku(Buku1)
    tampilBuku(Buku2)
    
}

func tampilBuku(buku Buku){
    fmt.Printf("judul buku %s\n", buku.judul)
    fmt.Printf("penulis %s\n", buku.penulis)
    fmt.Printf("id buku %d\n", buku.id_buku)
}
```
