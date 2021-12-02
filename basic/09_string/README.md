# string

string yang banyak digunakan dalam pemograman go, adalah sepotong byute yang hanya dapat dibaca. dalam bahasa pemgoraman go, string adalah **irisan**. platform go mentediakan bebagai perpustakaan untuk memanipulasi string

- unicode
- regexp
- strings

## membuat string

cara paling langsung untuk embuat string adalah dengan menulis

```golang
var pesan = "bellsahde"
```

setiap kali menemukan string literal dalam kode anda. kompilator membuat objek string dengan nilainya.literal string menyimpan urutan ``utf-8`` yang balid disebut rune. sebuah string memegang byte.

```golang
package main

import "fmt"

func main(){
    var pesan = "bellshade"

    fmt.Printf("normal string ")
    fmt.Printf("%s", pesan)
    fmt.Printf("\n")
    fmt.Pritnf("hex byte: ")

    for i := 0; < len(pesan); i++{
        fmt.Printf("%x", pesan[i])
    }
    
    fmt.Printf("\n")
    const contohPesan = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    fmt.Printf("string menggunakan quote ")
    fmt.Printf("%+q", contohPesan)
    fmt.Printf("\n")
}
```

## panjang string

``len(str)`` mengembalikan jumlah byte yang terkandung dalam string literal

```golang
package main

import "fmt"

func main(){
    var pesan = "bellshade"

    fmt,Printf("panjang string ")
    fmt.Println(len(pesan))
}
```

## menggabungkan string

packet sting memyertakan metode **join** untuk menggabungkan beberapa string
```
string.Join(sample, "")
```

join menggabungkan elemen-elemen array untuk membuat string tunggal. parameter kedua adalah pemisah yang ditempatkan di antara element array

```golang
package main

import ("fmt" "math") "fmt" "string")

func main(){
    pesan := []string{"bell", "shade"}
    fmt.Println(strings.Join(pesan, " "))
}
```