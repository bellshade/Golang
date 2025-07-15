## Embed

`embed` adalah standard library bawaan Go yang digunakan untuk menyisipkan file atau folder statis (seperti teks, gambar, konfigurasi, dll) langsung ke dalam binary Go saat compile time. Ini sangat berguna untuk menyertakan file seperti template HTML, konfigurasi, skrip, atau file lain yang dibutuhkan aplikasi tanpa harus membaca dari sistem file saat runtime.

Library `embed` mendukung penyisipan:

- File tunggal sebagai `string` atau `[]byte`
- Folder beserta isinya sebagai `embed.FS` (filesystem virtual)

### Tujuan

- Menyertakan file seperti template HTML, konfigurasi, aset, dll.
- Membuat aplikasi mandiri tanpa ketergantungan luar
- Meningkatkan performa dengan menghindari pembacaan dari disk saat runtime

### Contoh Kasus

#### Variabel Tunggal

Membaca isi file single_file.txt sebagai string

```go
//go:embed folder/single_file.txt
var fileString string
```

Membaca isi file single_file.txt sebagai slice of byte.

```go
//go:embed folder/single_file.txt
var fileByte []byte
```

#### Folder dan Multiple File

Membaca file single_file.txt dan semua file dengan ekstensi .hash dalam folder folder/ sebagai filesystem virtual.

```go
//go:embed folder/single_file.txt
//go:embed folder/\*.hash
var folder embed.FS
```

### Contoh Implementasi

```go
package main

import (
    "embed"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
    // Mencetak isi fileString
    println(fileString)

    // Mencetak isi fileByte sebagai string
    println(string(fileByte))

    // Membaca file1.hash dari folder
    content1, _ := folder.ReadFile("folder/file1.hash")
    println(string(content1))

    // Membaca file2.hash dari folder
    content2, _ := folder.ReadFile("folder/file2.hash")
    println(string(content2))
}
```
