# float

tipe data float digunakan untuk menyimpan nilai dengan titik desimalnya ``6.7``, ``-88.21`` kata kunci digunakan yaitu ``float32`` dan ``float64``

```go
var hargaDiskon float64
```

ada dua ukuran data float dalam pemograman golang yaitu ``float32`` yaitu yang memikili ukuran 4 byte pada arsitektur 32bit. ``float64`` yaitu  yang memiliki ukuran 8 byte pada arsitektur 64bit.

jika kita mendifiniikan variabel float tanpa menentukan ukuran secara eskplisit, ukuran variabel akan menjadi 64bit

```go
hargaDiskon := 521.2
```

```go
package main

import "fmt"

func main() {
    var hargaDiskon float32
    var hargaDiskon2 float64

    hargaDiskon = 5000.2141

    // meyimpan data yang lebih banyak
    hargaDiskon2 = 76124.9257571245

    fmt.Println(hargaDiskon);

    fmt.Println(hargaDiskon2);
}
```
