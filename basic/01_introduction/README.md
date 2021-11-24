# Pengenalan

- [Gambaran Singkat](#gambaran-singkat)
- [Persiapan](#persiapan)
  - [Instalasi Go](#instalasi-go)
  - [Cara Menjalankan Program](#cara-menjalankan-program)
- [Struktur Program Go](#struktur-program-go)
  - [Package](#package)
  - [Import](#import)
  - [Fungsi `main()`](#fungsi-main)
- [Menulis Program Go Pertamamu](#menulis-program-go-pertamamu)

## Gambaran Singkat

Go (atau biasa disebut Golang) merupakan bahasa pemrograman yang dikembangkan oleh Google dan dibuat dari bahasa C. Karena itu, banyak kemiripan dengan bahasa tersebut. Bagi kamu yang sudah pernah mempelajari C, Golang pasti akan terasa familier. Golang mulai dikenalkan ke publik pada tahun 2009. Jadi, bahasa ini masih terbilang *baru* untuk sebuah bahasa pemrograman.

Struktur dari Golang adalah sebagai berikut:

```go
package main

func main() {
  // kode kamu berada di sini
}
```

## Persiapan

### Instalasi Go

Sebelum mulai belajar, pastikan kamu sudah menginstall Go di komputer kamu. Jika belum, kamu bisa mengikuti panduan berikut ini.

1. Download *installer* dari [https://golang.org/doc/install](https://golang.org/doc/install). Pilih sesuai dengan OS yang kamu gunakan.
2. Setelah ter-*download*, jalankan installer dengan cara sebagai berikut.
    
    #### Windows
    
    - Buka file yang sudah kamu download lalu ikuti proses instalasi sampai selesai. Secara default, Go akan terinstall di `C:\go`.
    - Buka *command prompt* (CMD) **kamu dan ketikkan
        
        ```powershell
        > go version
        ```
        
    - Pastikan keluar output berupa versi Go yang sudah kamu install. Jika tidak, restart CMD kamu lalu coba kembali command di atas,
    
    #### Mac
    
    - Buka file yang sudah kamu download lalu ikuti proses instalasi sampai selesai. Secara default, Go akan terinstall di `/usr/local/go`.
    - Buka terminal kamu dan ketikkan
        
        ```bash
        $ go version
        ```
        
    - Pastikan keluar output berupa versi Go yang sudah kamu install. Jika tidak, restart CMD kamu lalu coba kembali perintah di atas,
    
    #### Linux
    
    - Buka terminal dan ekstrak file yang sudah didownload ke `/usr/local`. Kamu bisa menjalankan perintah berikut
        
        ```bash
        $ sudo tar -C /usr/local -xzf <nama file>
        ```
        
    - Tambahkan path instalasi Go ke dalam environment variable dengan cara
        
        ```bash
        $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
        $ source ~/.profile
        ```
        
    - Pastikan Go sudah terinstall dengan menjalankan perintah ini
        
        ```bash
        $ go version
        ```
        

### Cara Menjalankan Program

Program Go bisa dijalankan dengan dua cara

1. Melakukan kompilasi (*compile*) terlebih dahulu, lalu mengeksekusi file hasil kompilasi tersebut.
    
    ```bash
    $ go build main.go
    $ ./main
    ```
    
2. Langsung menjalankan program tanpa melakukan kompilasi
    
    ```bash
    $ go run main.go
    ```
    

## Struktur Program Go

Struktur dari bahasa Go adalah sebagai berikut:

```go
package main

import (
  // packages yang kamu gunakan
)

func main() {
  // kode kamu berada di sini
}

```

Mari kita bedah masing-masing keyword pada struktur tersebut.

### Package

```go
package <nama_package>
package main
```

Setiap program Go terbuat dari paket-paket (*packages*). Keyword `package` berfungsi untuk mengelompokkan code dalam sebuah project. Semua program Go harus diawali dengan nama package. Package `main` adalah yang pertama kali dieksekusi ketika menjalankan program.

### Import

```go
import <nama_package>
import "fmt"
import (
  "fmt"
  "os"
)
```

Dalam suatu program, jika kamu membutuhkan sesuatu dari package lain, kamu bisa memanggilnya dengan mengimport package tersebut. Dengan mengimport package `fmt`, kamu bisa mengakses fungsi-fungsi di dalam package tersebut, misalnya `fmt.Println()` dan `fmt.Scanf()`.

### Fungsi `main()`

```go
func main() {
  // kode kamu berada di sini
}
```

Dalam sebuah program Go, harus terdapat package `main` dan fungsi `main()` di dalamnya. Fungsi `main()` adalah fungsi yang menampung semua program kamu. Ketika program dijalankan, fungsi `main` akan dieksekusi pertama kali.

> **Tambahan:** Penggunaan titik koma (*semicolon*) tidak wajib (opsional)
> 

## Menulis Program Go Pertamamu

Program Go disimpan dengan ekstensi `.go`, Silakan buka code editor favorit kamu lalu tulis code di bawah ini dan simpan filenya dengan nama `hello_world.go`.

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

Penjelasan tentang code ini dapat dilihat [di sini](hello_world.go) 

Setelah itu, jalankan program menggunakan perintah

```bash
go run hello_world.go
```

Selamat! Kamu sudah berhasil menulis program pertamamu! Happy learning see you next lesson!

**Catatan dalam penamaan file:**

- Disarankan untuk memberi nama file tanpa spasi
- Gunakan `_` atau `-` jika nama file lebih dari satu kata, misal `hello_world.cpp` atau `hello-world.cpp`