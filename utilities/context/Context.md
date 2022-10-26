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

## Membatalkan Context (Context Cancellation)
Pembatalan (*cancellation*) berguna untuk mencegah sistem melakukan hal yang tidak perlu. Perhatikan HTTP request di bawah ini (gambar diambil dari [https://www.sohamkamani.com/golang/context-cancellation-and-values/#context-cancellation-in-go](https://www.sohamkamani.com/golang/context-cancellation-and-values/#context-cancellation-in-go)).

![context-cancellation-example](assets\context-cancellation-example.svg)

Pada gambar di atas adalah gambaran peristiwa yang terjadi jika tidak ada *context cancellation* pada HTTP request. Bisa kita lihat, aplikasi akan tetap memproses request dan melakukan query pada database **walaupun** user telah membatalkan request ke HTTP server. Kita bisa membatalkan seluruh proses pada aplikasi HTTP server kita dengan menggunakan *context cancellation* sehingga diagram di atas akan menjadi seperti ini.

![context-cancellation-example-2](assets\context-cancellation-example-2.svg)

Gambar di atas adalah skenario ketika kita melakukan *context cancellation* pada user request, dan seluruh proses yang bergantung pada context user request akan dibatalkan.

Di dalam interface package `context`, terdapat interface method sebagai berikut (kode diambil dari [https://www.prakharsrivastav.com/posts/golang-context-and-cancellation/](https://www.prakharsrivastav.com/posts/golang-context-and-cancellation/)).

```Go
type Context interface {
  Deadline()(deadline time.Time, ok bool)  // time when the ctx will cancel
  Done() <-chan struct{}                   // channel to track ctx cancellation
  Err() error                              // reason for ctx cancellation
  Value(key interface{}) interface{}       // to pass request-scoped values
}
```

Tiga method teratas digunakan untuk melakukan *context cancellation*. 

Berikut ini adalah cara yang bisa kita lakukan untuk membuat *context cancellation*.

1. *Listening* terhadap *context cancellation*
Untuk melakukan *listen* terhadap *context cancellation*, kita harus menunggu ada data yang dikirm dari channel `.Done()`. Perhatikan kode di bawah ini (kode diambil dari [https://www.sohamkamani.com/golang/context-cancellation-and-values/#context-cancellation-in-go](https://www.sohamkamani.com/golang/context-cancellation-and-values/#context-cancellation-in-go))

```Go
func main() {
	// Create an HTTP server that listens on port 8000
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Fprint(os.Stdout, "processing request\n")
		// We use `select` to execute a piece of code depending on which
		// channel receives a message first
		select {
		case <-time.After(2 * time.Second):
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// We then write this as the response
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request gets cancelled, log it
			// to STDERR
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))
}
```

Jika kita jalankan kode di atas dan membuka `localhost:8000` kurang dari 2 detik, maka akan terjadi *context cancellation* di mana blok kode `<-ctx.Done()` tereksekusi.

2. Membuat event *context cancellation*
Jika kita menghendaki ada proses yang di-*cancel*, kita harus membuat event yang menghasilkan *context cancellation*. Hal ini bisa lakukan dengan menggunakan function `context.WithCancel(ctx)` yang akan mengembalikan `context` dan *cancellation function* `. Perhatikan kode di bawah ini (kode diambil dari (https://www.sohamkamani.com/golang/context-cancellation-and-values/#context-cancellation-in-go)[https://www.sohamkamani.com/golang/context-cancellation-and-values/#context-cancellation-in-go])

```Go
func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func main() {
	// Create a new context
	ctx := context.Background()
	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)

	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}
```

Ketika kita menjalankan kode di atas, kita akan mendapatkan hasil `halted operation2`. Hal disebabkan karena kita **sengaja** melakukan *context cancellation* ketika kita mendeklarasikan error pada `err := operation1(ctx)`, di mana error tersebut sengaja dihasilkan supaya kita bisa men-*trigger* cancellation function, dalam hal ini cancellation function yang dimaksud adalah function `cancel()` pada blok `if-else`. Karena cancellation function dijalankan, maka `context` yang digunakan pada function `operation2` akan menerima *cancellation signal* dari `ctx.Done()`, sehingga kita mendaptkan hasil `halted operation2` di console kita.

Sumber : 
- https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
- https://www.sohamkamani.com/golang/context-cancellation-and-values/#context-cancellation-in-go
- https://www.prakharsrivastav.com/posts/golang-context-and-cancellation/

