# Byte Golang

Dalam bahasa Go, **`byte`** adalah tipe data yang sangat penting untuk menangani **data biner** dan **karakter berbasis ASCII**. `byte` merupakan **alias dari `uint8`**, yang artinya ia adalah **bilangan bulat 8-bit tanpa tanda**, dengan nilai dari **0 hingga 255**.
`byte` biasa digunakan dalam pengolahan teks ASCII, buffer data, komunikasi jaringan, encoding, serta pembacaan/penulisan file.

---

## **Byte Literal (Karakter ASCII)**

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var b byte = 'A' // Karakter ASCII untuk A = 65

    // menampilkan byte dan tipe nya
    fmt.Printf("byte: %c; nilai: %d; tipe: %s\n", b, b, reflect.TypeOf(b))
}
```

## **Byte Literal**

- Literal `byte` dapat didefinisikan dari:
  - **Angka**: `var b byte = 65`
  - **Karakter ASCII**: `var b byte = 'A'`
- Tipe data byte **tidak bisa menggunakan Unicode kompleks atau emoji** seperti `'你'` atau `'😊'` karena nilainya di luar jangkauan `byte`.

---

## **Byte dalam Slice (`[]byte`)**

Go menggunakan slice of bytes (`[]byte`) untuk representasi data teks dan file:

```go
package main

import "fmt"

func main() {
    data := []byte("Hello") // Konversi string ke byte slice
    fmt.Println(data)       // Output: [72 101 108 108 111]
}
```

---

## **Byte vs Rune**

| Aspek         | `byte` (`uint8`)            | `rune` (`int32`)                |
| ------------- | --------------------------- | ------------------------------- |
| Ukuran        | 8-bit                       | 32-bit                          |
| Rentang Nilai | 0 – 255                     | -2.1 miliar – +2.1 miliar       |
| Representasi  | Karakter ASCII / Data biner | Karakter Unicode                |
| Kegunaan Umum | File, jaringan, ASCII       | Teks multibahasa, simbol, emoji |
| Slice Umum    | `[]byte`                    | `[]rune`                        |
