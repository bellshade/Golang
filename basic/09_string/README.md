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
```go
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

## memecah string menjadi slice rune

Dalam bahasa Go, kita bisa memecah sebuah string menjadi kumpulan rune dalam sebuah slice. Misal kita punya string `"Hello"`, maka akan diconvert jadi sebuah slice `['H','e','l','l','o']`.

Syntax: `strSlice := []rune(variableString)`

Contoh:

```go
str := "Hello"

strSlice := []rune(str)

fmt.Println(strSlice)
```
Output:

`[72 101 108 108 111]`

Jangan bingung dengan output angka-angka diatas. Itu hanya ASCII Code (representasi angka dari tiap karakter).

Hint: Teknik `string to rune slice` ini cocok diterapin kalo kita mau sorting sebuah string