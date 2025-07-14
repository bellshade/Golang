## net/http

`net/http` adalah standard library bawaan Go yang digunakan untuk membangun aplikasi berbasis protokol HTTP. Tujuan antara lain:
- Menangani request dari client seperti API atau browser
- Mengirim request ke server lain
- Middleware, routing dasar, header manipulation, cookie dan lain-lain.

Ada dua bagian utama dalam `net/http`:

1. Server Side
Membuat server yang akan listen request dari client, yang terdiri dari
- `http.Request` yang merepresentasikan request HTTP yang masuk
- `http.ResponseWriter` yang nantinya akan digunakan dalam memberikan response ke client
- `http.Handler` interface untuk menangani request dan response
- `http.HandleFunc` untuk register handler yang berdasarkan dari path URL

2. Client Side
Digunakan untuk melakukan request ke server lain (contoh: `GET/POS` ke `API` eksternal), dalam component client side terdapat
- `http.Get(url string)`, `http.Post(...)`
- `http.Client{}` digunakan untuk konfigurasi custom request (seperti request, headers, dan lain-lain)

## contoh implementasi

Fungsi `homeHandler` yang akan dipanggil saat user mengakses `/`. dan mencatat response ke client.
```go
func homeHandler(rw http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(rw, "selamat datang di bellshade Golang !\n")
}
```

Fungsi `aboutHandler` yang akan dipanggil sama seperti `homeHandler` tetapi bedanya dia digunakan untuk page lain misalnya `/about`
```go
func aboutHandler(rw http.ResponseWriter, req *http.Request) {
    fmt.Printf(rw, "kita di page about - bellshade Golang\n")
}
```

Fungsi `loggingMiddleware` adalah fungsi wrapping handler lain dan akan record log setiap request-request yang akan masuk
```go
func loggingMiddleware(hf http.HandlerFunc) http.HandlerFunc {
    return func(rw http.ResponseWriter, req *http.Request) {
        fmt.Printf("request sudah diterima: %s %s\n", rw.Method, req.URL.Path)
        hf.ServeHTTP(rw, req)
    }
}
```

kita bisa mencoba dengan menggunakan fungsi `http.HandleFunc(...)` untuk menghubungkan URL path ke handler

```go
http.HandleFunc("/", loggingMiddleware(homeHandler))
http.HandleFunc("/about", loggingMiddleware(aboutHandler))

fmt.Println("server berjalan dengan http://localhost:8443")
```
ketika dijalankan output akan
```
server berjalan dengan http://localhost:8443
```
