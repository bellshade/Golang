# struct

golang tidak mendukung kustom melalui kelas tetapi _struct_. struct di golang adalah tipe yang ditentukan pengguna yang dapat hnaya menyimpan status dn bukan perilaku. struct dapt digunakan untuk mempresentasikan objek kompleks yang terdiri dri lebih satu pasangan nilai kunci. kita dapat menambahkan fungsi ke struct yang dapat menambahkan perlaku seperti contoh.

```go
package main

import "fmt"

// deklarasi struct
type Siswa struct {
  nama string
  kelas int
}

// fungsi untuk menampilkan detail dri siswa
func (siswa Siswa) print_details() {
  fmt.Printf("nama siswa; %s", siswa.nama)
  fmt.Printf("kelas: %d", siswa.kelas)
}

func main() {
  // deklarasi struct
  siswa1 := Siswa{"James", 4}
  
  // print dri detail
  siswa1.print_details()

  siswa.nama = "sundar"
  siswa.kelas = 6
}
```
