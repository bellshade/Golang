# Percabangan

- [Percabangan](#percabangan)
  - [Overview](#overview)
  - [1. Percabangan dengan **if else**](#1-percabangan-dengan-if-else)
    - [A. Penggunaan **if else** Dasar](#a-penggunaan-if-else-dasar)
    - [B. Penggunaan Temporary Variable](#b-penggunaan-temporary-variable)
  - [2. Percabangan dengan **switch case**](#2-percabangan-dengan-switch-case)
    - [A. Penggunaan **switch case** Dasar](#a-penggunaan-switch-case-dasar)
    - [B. Penggunaan seperti **if else**](#b-penggunaan-seperti-if-else)

## Overview

**Condition/Percabangan** merupakan mekanisme dalam pemrograman yang bertujuan untuk membatasi agar sebuah blok kode hanya berjalan jika kondisi yang diberikan bernilai *true*.

Dalam bahasa pemrograman Golang, terdapat 2 tipe percabangan, yaitu dengan menggunakan **if else** dan **switch**.

## 1. Percabangan dengan **if else**

### A. Penggunaan **if else** Dasar

Penggunaan percabangan **if else** di Golang tidak jauh berbeda dengan bahasa pemrograman yang lainnya. Jika dalam bahasa pemrograman lain penulisan *condition statement* harus diawali dan diakhiri dengan tanda kurung, maka di Golang kita tidak perlu melakukan hal tersebut. Agar lebih mudah, silakan lihat pada blok kode di bawah ini.

```golang
numChild := 5

if numChild > 2 {
    fmt.Println("Banyak anak banyak rejeki")
} else {
    fmt.Println("Dua anak cukup")
}
```

### B. Penggunaan Temporary Variable

Selain itu juga, percabangan **if else** yang ada di Golang dapat menggunakan *temporary variable* yang mungkin di bahasa pemrograman lain hal itu belum tersedia. *Temporary Variable* yang dimaksud adalah sebuah variabel yang hanya digunakan pada satu kode blok tersebut saja. Agar lebih mudah dalam memahami, silakan perhatikan blok kode di bawah ini.

```golang
if numChild := 5; numChild > 3 {
    fmt.Println("Banyak anak banyak rejeki")
} else {
    fmt.Println("Cukup aja")
}
```

## 2. Percabangan dengan **switch case**

### A. Penggunaan **switch case** Dasar

Penggunaan percabangan dengan **switch case** dalam bahasa pemrograman Golang tidak jauh berbeda dengan yang ada dalam bahasa pemrograman lain. Sama dengan percabangan **if else**, penulisan *condition statement* tidak perlu menggunakan tanda kurung. Lihat kode blok di bawah ini agar lebih mudah dimengerti.

```golang
numChild := 5

switch numChild {
case 5:
    fmt.Println("Mantap")
case 3:
    fmt.Println("Bagus")
default:
    fmt.Println("Semangat")
}
```

### B. Penggunaan seperti **if else**

Selain itu juga, percabangan dengan **switch case** bisa juga digunakan layaknya seperti memakai percabangan **if else**, caranya adalah sebagai berikut.

```golang
numChild := 5

switch {
case numChild > 5:
    fmt.Println("Lebih dari lima")
case numChild > 3:
    fmt.Println("Lebih dari tiga")
default:
    fmt.Println("Kurang dari tiga")
}
```
