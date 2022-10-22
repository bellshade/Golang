# Context

## Definisi
Ketika kita mengembangkan aplikasi web server, seringkali kita hanya perlu mengetahui URL yang dikirimkan ke HTTP request tanpa mempedulikan apakah user masih tersambung ke server atau tidak. Jika user tiba-tiba tidak tersambung (*disconnected*) sebelum user menerima response, server akan tetap memproses response hingga user tersambung kembali. Namun, proses tersebut tentunya akan memakan sumber daya server yang tentunya bisa digunakan untuk meng-*handle* tugas lain daripada menununggu user tersambung kembali kemudian mengirimkan response ke user tersebut.

Dalam hal ini, mengetahui konteks (selanjutnya disebut *context*) menjadi hal yang penting, seperti apakah user masih tersambung dengan server atau tidak. Dengan mengetahui hal tersebut, kita bisa menghemat sumber daya komputasi dan server bisa meng-*handle* request dari user lain.

Di `Golang`, package `context` digunakan untuk menangani hal tersebut. `Context` bisa kita gunakan untuk mengetahui proses yang berjalan pada background hingga memberi tanda untuk berhenti (*termination signal*).

## Membuat Context
Banyak package di `Golang` yang menggunakan package `context` untuk mengolah informasi tambahan dari *environment* yang sedang dieksekusi. Dengan menggunakan package `context`, kita bisa mengetahui proses yang terjadi dari suatu function ke function lain, dari awal hingga akhir. `Context` di `Golang` merupakan package bawaan, sehingga kita tidak perlu meng-*install* menggunakan `go get`. Berikut cara menggunakan `context` di `Go`

```Go
package main

import (
	"context"
	"fmt"
)

func doSomethingWithContext(ctx context.Context){
	fmt.Printf("This is : %v \n",ctx)
}

func main() {
	ctx_todo := context.TODO()
	ctx_background := context.Background()
	// menghasilkan This is : context.TODO
	doSomethingWithContext(ctx_todo)
	// menghasilkan This is : context.Background
	doSomethingWithContext(ctx_background)
}
```

*By convention*, `context` di `Golang` dideklarasikan dengan `ctx`. Dari kode di atas, kita mendeklarasikan `context.TODO()` dan `context.Background()`. Jika kita tidak tahu `context` apa yang harus digunakan, kita bisa menggunakan `context.TODO()` atau `context.Background()`. Kedua `context` tersebut mengembalikan `context` kosong. Namun, biasanya `context.Background()` digunakan ketika developer belum tahu `context` mana yang harus digunakan.

## Mengirim Data Dari Satu Context Ke Context Lain
Salah satu keuntungan `context` adalah kita bisa menyimpan data dalam `context`. Misalkan suatu function membutuhkan username dan password, dan function lain membutuhkan filepath dari user, dan function terakhir akan mengakses file yang dikirimkan function sebelumnya. Hal tersebut bisa dilakukan dengan menyimpan data ke dalam `context`.

Di `Golang`, kita bisa menggunakan `context.WithValue` untuk menambahkan data ke `context`. Berikut contohnya

```Go
package main

import (
	"context"
	"fmt"
)

func doSomethingContextWithValue(ctx context.Context){
	fmt.Printf("key from ctx_with_val : %s \n", ctx.Value("key"))
}

func main() {
	ctx_background := context.Background()
	ctx_with_val := context.WithValue(ctx_background, "key", "val")
	// menghasilkan key from ctx_with_val : val
	doSomethingContextWithValue(ctx_with_val)
}
```

Seperti yang sudah dijelaskan tadi, kita bisa mengirimkan nilai `context` dari satu function ke function lain. Perhatikan contoh berikut.

```Go
func doAnotherFromSomething(ctx context.Context){
    // menghasilkan anotherKey from doAnotherCtx anotherKey : anotherVal
	fmt.Printf("anotherKey from doAnotherCtx anotherKey : %s \n", ctx.Value("anotherKey"))
}

func doSomethingContextWithValue(ctx context.Context){
	fmt.Printf("key from ctxWithVal : %s \n", ctx.Value("key"))
	doAnotherCtx := context.WithValue(ctx, "anotherKey", "anotherVal")
	doAnotherFromSomething(doAnotherCtx)
	fmt.Printf("key from ctxWithVal : %s \n", ctx.Value("key"))
}

func main() {
	ctxWithVal := context.WithValue(ctxBackground, "key", "val")
	// menghasilkan key from ctx_with_val : val
	doSomethingContextWithValue(ctxWithVal)
    // hasil akhir
	// key from ctxWithVal : val
	// anotherKey from doAnotherCtx anotherKey : anotherVal
	// key from ctxWithVal : val
}
```
Setelah kita jalankan program di atas, kita akan mendapatkan hasil sebagai berikut
```
key from ctxWithVal : val
key from doAnotherCtx anotherKey : anotherVal
key from ctxWithVal : val
```
Baris pertama adalah `context` yang dihasilkan oleh function `doSomethingContextWithValue` baris kedua dari function `doAnotherFromSomething` dan baris ketiga dari function `doSomethingContextWithValue` lagi. Seperti yang kita lihat juga, bahwa nilai dari `context` dari `doSomethingContextWithValue` ke `doAnotherFromSomething` tidak berubah (tidak di-*override*), karena `context` hanya akan me-*wrapping* `key` `context` terluar, dalam hal ini `doAnotherCtx` adalah `context` terluar.

Sumber : https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

