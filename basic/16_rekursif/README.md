# Rekursif

rekursi adalah proses pengulangan item dengan cara yang serupa. konsep yang sama juga berlaku dalam bahasa pemograman. jika sebuah program memungkinkan untuk memanggil fungsi di dalam fungsi yang sama, maka itu disebut pemanggilan fungsi rekursif. contoh

```golang
func rekursif(){
  /* fungsi memanggil dirinya sendiri */
  rekursif()
}

func main(){
  rekursif()
}
```

bahasa pemograman **Go** mendukung rekursi. artinya, memungkinkan suatu fungsi untuk memanggil dirinya sendiri. tetapi saat menggunakan rekursi, pemogram harus berhati-hati untuk mementukan kondisi keluar dari fungsi, jika tidak maka akan menjadi loop tidak terbatas.

## contoh rekursi di Go

fungsi rekursif sangat berguna untuk menyelesaikan banyak masalah matematika seperti menghitung faktorial suatu bilangan, menghasilkan deret _fibonacci_, dll.

**contoh 1: menghitung faktorial menggunakan rekursi di go**

```golang
package main

import "fmt"

func factorial(i int)int {
  if (i <= 1) {
      return 1
  }
  return i * factorial(i - 1)
}

func main() {
  var i int = 10
  fmt.Printf("Factorial dari %d adalah %d", i, factorial(i))
}
```

**contoh 2: menghitung deret fibonnaci menggunakan rekursi di go**

```golang
package main

import "fmt"

func fibonnaci(i int) (ret int) {
  if i == 0 {
    return 0
  }
  if i == 1 {
    return 1
  }

  return fibonnaci(i - 1) + fibonnaci(i - 2)
}

func main() {
  var i int
  for i = 0; i < 10; i++ {
    fmt.Printf("%d", fibonnaci(i))
  }
}
```

