# Goroutine

Goroutine dapat dianggap sebagai thread ringan yang memiliki eksekusi independen yang terpisah dan yang dapat dieksekusi secara bersamaan dengan goroutine lainnya. Ini adalah fungsi atau metode yang dieksekusi secara bersamaan dengan goroutine lain. Ini sepenuhnya dikelola oleh runtime GO. Golang adalah bahasa konkuren. Setiap goroutine adalah eksekusi independen. Ini adalah goroutine yang membantu mencapai konkurensi di golang

Golang menggunakan kata kunci khusus '__go__' untuk memulai goroutine. Untuk memulai, cukup tambahan kata kunci **go** sebelum pemanggilan fungsi atau metode. Fungsi atau metode itu sekarang akan dieksekusi di goroutine. Perhatikan bahwa bukan fungsi atau metode yang mementukan apakah itu goroutine. jika kita memanggil metode atau fungsi tersebut dengan kata kunci go, maka fungsi atau metode tersebut dikatakan dieksekusi dalam goroutine.

Mari kita pahami perbedaan antara menjalankan fungsi secara normal dan menjalankan fungsi sebagai goroutine.

```
perintah
start()
perintah2
```

dalam menjalankan fungsi normal untuk skenario di atas

- pertama, ``perintah`` akan dieksekusi
- kemudian fungsi ``start()`` akan dipanggil
- setelah fungsi ``start()`` selesai maka ``perintah2`` akan dieksekusi


**menjalankan fungsi goroutine**

```
perintah
go start()
perintah2
```

- pertama, ``perintah`` akan dieksekusi
- kemudian fungsi ``start()`` akan dipanggil sebagai goroutine yang akan dieksekusi secara asynchronous.
- ``perintah2`` akan seegera dieksekusi. Itu tidak akan menunggu fungsi ``start()`` selesai. Fungsi start akan dieksekui secara bersamaan sebagai goroutine sementara sisa program melanjutkan eksekusinya.

Jadi pada dasarnya ketika memanggil suatu fungsi sebagai goroutine, panggilan akan segera kembali, eksekusi akan dilanjutkan dari baris berikutnya sementara goroutine akan dieksekusi secara bersamaan di latar belakang. Perhatikan juga bahwa setiap nilai kembalian dari goroutine akan diabaikan.


**contoh**
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go start()
    fmt.Println("bellshade")
    time.sleep(1 * time.Second)
    fmt.Println("selesai!")
}

func start() {
    fmt.Println("menjalankan fungsi goroutine")
}
```

dalam program di atas kita menggunakan kata kunci '__go__' sebelum pemanggilan fungi untuk memulai goroutine.

```
go start()
```

baris diatas akan memulai goroutine yang akan menjalankan fungsi ``start()``. Program pertama kali mencetak "__bellshade__". Perhatikan bahwa baris di mana kita mencetak "__bellshade__" adalah setelah goroutine dimulai. Ini menggambarkan poin yang disebutkan di atas bahwa setelah goroutine dimulai, panggilan dilanjutkan dari baris selanjutnya. kita kemuidan memberikan batas waktu. Batas waktu ada sehingga goroutine terjadwal sebelum goroutine utama ada. jadi sekarang goroutine dijalankan dan dicetak

```
menjalankan fungsi goroutine
```

kemudian dicetak
```
selesai!
```

```go
package main
import (
    "fmt"
)

func main() {
    go start()
    fmt.Println("bellshade")
    fmt.Println("selesai")
}

func start() {
    fmt.Println("menjalankan fungsi goroutine")
}
```