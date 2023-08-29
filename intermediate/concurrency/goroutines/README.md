# Councurrency
```
Concurrency adalah proses menjalankan beberapa komputasi secara bersamaan.
```

Analogi sederhana untuk memahami konsep *concurrency* adalah proses memasak. Misalkan, seseorang sedang memasak mi instan dengan tambahan *topping* berupa telur rebus dan sosis goreng di dalam satu kompor yang sama. Pertama, ia memanaskan air terlebih dahulu untuk merebus telur. Selagi menunggu air tersebut mendidih, ia memanaskan wajan untuk menggoreng sosis. Sembari ia menunggu wajan tersebut panas, ia membuka bungkus mi instan dan mengambil piring untuk menyajikan mi tersebut jika semua bahan sudah matang. Ternyata, wajan yang disiapkan tadi sudah siap digunakan untuk menggoreng sosis dan ia langsung menggoreng sosis. Selagi ia menunggu sosis matang, ia merebus telur karena airnya sudah mendidih. Ia kembali ke kompor yang digunakan untuk menggoreng sosis dan ia meniriskan sosis tersebut. Kemudian ia kembali ke kompor yang ia gunakan untuk merebus telur dan ia meniris telur tersebut. Tidak lupa, ia memasak mi instan dan menunggu mi tersebut matang. Terakhir, ia kumpulkan seluruh bahan dalam piring yang telah ia siapkan di awal.

Proses memasak di atas divisualisasikan pada diagram berikut

![councrrency_example](assets\concucrrency_example.png)

Pada diagram di atas, ada beberapa proses yang diwakilkan oleh persegi panjang dengan warna *border*, yaitu proses yang berhubungan dengan Telur, Sosis, dan Mi. Dalam *concurrency*, proses-proses tersebut dikenal dengan istilah *routine*. Seluruh proses memasak tersebut dilakukan dalam satu kompor yang sama dengan asumsi kompor yang digunakan memiliki tempat memasak yang banyak. Dalam *concurrency*, tempat memasak pada kompor tersebut dikenal dengan istilah *thread*
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

```go
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
## Menjalankan Beberapa Goroutine Secara Bersamaan
Kita bisa menjalankan beberapa goroutine secara bersamaan. Perhatikan kode berikut (kode diambil dari [https://golangbot.com/goroutines/](https://golangbot.com/goroutines/))

```Go
package main

import (  
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
```

Ketika kita menjalankan kode di atas, kita akan mendapat *output* sebagai berikut.

`1 a 2 3 b 4 c 5 d e main terminated`

Visualisasi dari hasil di atas adalah sebagai berikut.

![councrrency_example2](assets\Goroutines-explained.png).

Seperti kita lihat pada gambar di atas, perintah `time.Sleep(3000 * time.Millisecond)` digunakan untuk menunggu goroutine `go alphabets()` dan `go numbers()` selesai dieksekusi. 

Selanjutnya, mari kita perhatikan goroutine `go alphabets()` dan `go numbers()`. Pada gambar di atas, bisa kita lihat bahwa kedua goroutine tersebut dijalankan secara bersamaan tanpa menunggu salah satu goroutine selesai dieksekusi (maksudnya selesai mencetak dari huruf a sampai e untuk goroutine `go alphabets()` dan mencetak angka 1 sampai 5 untuk goroutine `go numbers()`).

# Channel
**Channel** saluran dalam bentuk queue yang digunakan untuk berkomunikasi antar gorountine.  ***Secara default***, proses menerima dan mengirim data pada channel bersifat *blocking* atau saling menunggu. Jika ada data yang masuk ke dalam channel, data tersebut harus diterima terlebih dahulu, kemudian data yang baru bisa memasuki channel setelah data sebelumnya selesai diterima.

## Mendeklarasikan Channel
Berikut cara mendeklarasikan, mengirim dan menerima data dari channel (kode diambil dari [https://golangbot.com/channels/](https://golangbot.com/channels/))

```Go
a := make(chan int) // mendeklarasikan channel dengan tipe data integer

data := <- a // mengirim data dari channel a dan menyimpan hasilnya ke dalam variabel data
a <- data // channel a menerima data dari variabel data
```

## Contoh Penggunaan Channel
### Contoh Sederhana
Berikut ini adalah kode cara penggunaan channel (kode diambil dari https://golangbot.com/channels/)
```Go
package main

import (  
    "fmt"
)

func hello(done chan bool) {  
    fmt.Println("Hello world goroutine")
    done <- true
}
func main() {  
    done := make(chan bool)
    go hello(done)
    <-done
    fmt.Println("main function")
}
```

Hasil dari kode di atas adalah sebagai berikut

````
Hello world goroutine
main function
````

Pada program di atas, kita mendeklarasikan channel `done` dengan tipe data boolean. Kemudian kita memiliki goroutine `go hello()` yang argumennya kita isi dengan channel `done`. Karena pada dasarnya sifat channel adalah *blocking*, `main function` tidak akan ditampilkan ke layar sebelum channel `done` selesai menerima data.

### Memahami Lebih Dalam Konsep Blocking Pada Channel
Mari kita modifikasi kode di atas menjadi seperti ini (kode diambil dari [https://golangbot.com/channels/](https://golangbot.com/channels/)).

```Go
package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}

```

Kode di atas akan menghasilkan keluaran sebagai berikut.

```
Main going to call hello go goroutine
hello go routine is going to sleep
hello go routine awake and going to write to done
Main received data
```

Pada hasil keluaran di atas, kita akan mendapatkan `Main going to call hello go goroutine` ditampilkan pertama kali, kemudian `hello go routine is going to sleep` dicetak dari go routine `go hello(done)` dan program akan *sleep* selama 2 detik. Setelah waktu *sleep* berakhir, `hello go routine awake and going to write to done` dicetak, dan boolean `true` akan dikirim ke *channel* `done`. Terakhir, *channel* `done` menerima data dan `Main received data` (ingat bahwa `Main received data` tidak akan tercetak sebelum *channel* `done` selesai menerima data).

# Mutex
Ketika kita menggunakan `routines`/`goroutine`, kita bisa saja mendapati bahwa program yang kita buat ternyata terjadi *race condition*. ***Race condition*** adalah keadaan di mana banyak `routines` yang mengakses satu data yang sama pada saat yang bersamaan, sehingga ketika `routines` tersebut selesai mengakses/melakukan perubahan pada data tersebut, hasil akhir pada data tersebut tidak konsisten. Kode di bawah ini mendemonstrasikan kejadian *race condition*

```Go
func main() {
	x := 0

	for i := 1; i <= 1000; i++ {
		// will run 1000 go routines
		go func() {
			for j := 1; j <= 100; j++ {
				// x is share among 1000 go routines,
				// and should be returned 100000.
				x = x + 1
			}
		}()
	}

	time.Sleep(2 * time.Second)
	// each execution will give different result
	fmt.Printf("value of x : %v", x)
}
```

Setelah kita menjalankan kode di atas sebanyak 3 kali, kita akan mendapatkan hasil akhir dari variabel `x` yang berbeda setiap kali kita menjalankan kode tersebut, berikut hasilnya

```
value of x : 96671
value of x : 91098
value of x : 98686
```

Untuk mencegah terjadinya *race condition*, kita bisa menggunakan `Mutex` yang merupakan struct dari package bawaan `sync`.

## Implementasi Mutex
`Mutex` merupakan struct dari package bawaan `sync`. Pada kode di bawah ini, kita melakukan *lock* terhadap segala perubahan yang terjadi pada variabel `x`, sehingga tidak terjadi rebutan antar `goroutine` yang akan melakukan perubahan terhadap variabel `x`. Setelah seluruh perubahan selesai dikerjakan, kita melakukan *unlock* sehingga `goroutine` lain bisa melakukan perubahan terhadap variabel `x`.

```Go
func main() {
	x := 0
	var mutex sync.Mutex
	
	for i := 1; i <= 1000; i++ {
		// will run 1000 go routines
		go func() {
			for j := 1; j <= 100; j++ {
				// x is share among 1000 go routines,
				// and should be returned 100000.

				// locking any changes apply to x
				mutex.Lock()
				x = x + 1
				// unlocking the changes
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(2 * time.Second)
	// now each itteration will result the same
	fmt.Printf("value of x : %v", x)
}
```

Kita akan mendapatkan hasil yang sama setiap kali kita menjalankan kode di atas
```
value of x : 100000
value of x : 100000
value of x : 100000
```

# WaitGroup
```go
var wg sync.WaitGroup
```
 Apa itu **WaitGroup**, **WaitGroup** adalah fitur yang digunakan untuk menunggu satu atau beberapa proses goroutine selesai dieksekusi. 
 *contoh kode yang **tidak** mengunakan **Waitgroup***
 ```go
func sayHelloWorld(){
	fmt.Println("Hello World")
}
func main() {
	go sayHelloWorld()
	time.Sleep(1 * time.Second)
	fmt.Println("Kode tereksekusi")
}
```
Kode diatas sebenarnya hanya memperkirakan kapan waktu **goroutine** tersebut selesai, dan ini sangat **tidak dianjurkan**, karena dapat mengakibatkan masalah yang sulit diidentifikasi dan diatasi. Dalam contoh diatas, meskipun goroutine mungkin telah selesai dalam waktu satu detik, ada kemungkinan kasus di mana goroutine belum selesai, bahkan setelah waktu yang ditentukan. Oleh karena itu, disarankan untuk menggunakan mekanisme yang lebih andal seperti **WaitGroup** untuk mengontrol penyelesaian goroutine secara akurat.

----
deklarasi WaitGroup :
```go
var wg sync.WaitGroup
```
WaitGroup memiliki 3 method :
- `wg.Add(i int)` : adalah method untuk memberitahu `WaitGroup` tentang jumlah goroutine yang akan dijalankan
- `wg.Done()` : method yang digunakan di dalam setiap goroutine untuk memberi sinyal bahwa goroutine telah selesai.
- `wg.Wait()` : method yang digunakan untuk memblokir eksekusi program sampai semua goroutine yang ditambahkan selesai dan memanggil `Done()`.

*contoh kode yang menggunakan **WaitGroup*** :
```go
// If `WaitGroup` is used as a parameter, make sure the parameter is a pointer.
func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done() // Decreasing the count in the WaitGroup after the goroutine is finished
	fmt.Printf("Worker %d starting \n", x)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done \n", x)
}

  
func main() {
// Initialize WaitGroup
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1) // Adding 1 to the count of goroutines to be waited
		go worker(i, &wg)
	}
	wg.Wait() // Waiting until all goroutines are finished
	fmt.Println("All worker are finished")
}
```
mari bedah kode diatas satu persatu :
- **func worker(x int, wg *sync.WaitGroup)**: Ini adalah fungsi `worker` yang menerima dua parameter, yaitu `x` yang merupakan nomor identifikasi pekerja (goroutine), dan `wg` yang adalah pointer ke `sync.WaitGroup`. Penggunaan `*sync.WaitGroup` sebagai parameter menunjukkan bahwa parameter ini adalah referensi ke objek `WaitGroup` yang sama, yang akan digunakan untuk mengendalikan penyelesaian goroutine.
- **wg.Done()** : Pada akhir setiap goroutine, perintah ini mengurangkan nilai hitung di `WaitGroup` menggunakan metode `Done()`, menandakan bahwa goroutine tersebut telah selesai.
- **wg.Add(1)** dan **go worker(i, &wg)**: Dalam loop `for`, setiap iterasi menambahkan 1 ke hitungan `WaitGroup` dengan `wg.Add(1)` sebelum goroutine dimulai. Ini memastikan bahwa `WaitGroup` mengetahui bahwa ada goroutine yang harus selesai.
- **wg.Wait()**: Di dalam `main` function, `wg.Wait()` akan menunggu hingga semua goroutine yang telah dimulai selesai dieksekusi. Ini memastikan bahwa program tidak akan melanjutkan eksekusi lebih jauh sampai semua goroutine selesai.