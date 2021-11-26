# Operator

Operator adalah simbol yang memberitahu compiler untuk melakukan manipulasi natematika atau logika tertentu. bahasa golang kaya akan operator bawaan dan menyediakan jenis operator

- operator aritmatika
- operator rasional
- operator logika
- opeator bitwise
- operator penugasan
- opeator lain - lain

## operator aritmatika

tabel berikut menunjukkan semua operator aritmatika yang didukung oleh bahasa go

| operator      | keterangan                    | contoh        |
| :---          | :---                          | :---:         |
| +             | menambah dua operan           | A + B         |
| -             | mengurangu dua operan         | A - B         |
| *             | mengalikan dua operan         | A * B         |
| /             | membagi dua operan            | B / A         |
| %             | sisa bagi antara dua operan   | B % A         |
| ++            | operator penambahan satu integer  | A++       |
| --            | operator pengurangan satu integer | A--       |


## operator relasional

tabel berikut menjukkan semua operator relasional yang di dukung oleh bahasa go

| Operator      | keterangan                                    | contoh        |
| :---:         | :---                                          | :---:         |
| ``==``        | memeriksa apakah nilai dua operan sama        | ``a== b``     |
| ``!=``        | memeriksa apakah nilai dua operan tidak sama  | ``a!= b``     |
| ``>``         | memeriksa apakah nilai operan kiri lebih besar| ``a > b``     |
| ``<``         | memeriksa apakah nilai operan kanan lebih besar| ``a < b``    |
| ``>=``        | memeriksa apakah nilai operan kiri lebih besar atau sama dengan nilai operan kanan | `` a >= b``      |
| ``<=``        | memeriksa apakah nilai operan kanan lebih besar atau sama dengan nilai operan kanan | `` a <= b``     |


## operator logika

tabel berikut menujukkan semua operator logika yang di dukung oleh bahasa go

| Operator      | keterangan                                    | contoh        |
| :---:         | :---                                          | :---:         |
| ``&&``        | disebut operator logika AND. jike kedua operan salah, maka kondisi menjadi salah | ``A && B``         |
| ``||``        | disebut operator logika ATAU. jika salah satu operan bemar, maka kondisi menjadi benar | ``A || B``   |
| ``!``         | disebut operator BUKAN. gunakan untuk membalikkan keadaan logis operannya. jika suatu kondisi benar, maka operator logical not akan membuatnya salah | ``!(A && b)``        |


## operator bitwise

operator bitwise bekerja pada bit dan melakukan operasi bit demi bit. tabel untuk &, |, ^ adalah sebagai berikut

| var p      | var q      | var p & q         | var ``p | q``     | var p     |
| :---  | :---  | :---         | :---         | :--- |
| 0      | 0      | 0             | 0             | 0     |
| 0      | 1      | 0             | 1             | 1     |
| 1      | 1      | 1             | 1             | 0     |
| 1      | 0      | 0             | 1             | 1     |


asumsikan a = 60 dan b = 13. dalam format biner, mekera akan menjadi seperti berikut

```
a = 0011 1100
b = 0000 1101
---------------

a&b = 0000 1100
a|b = 0011 1101
a^b = 0011 0001
~A = 1100 0011
```

operator bitwise yang didukung oleh bahasa c tercantum dalam list berikut. asumsikan variabel A menampung 60 dan variabel B menampung 13


- ``&``

  biner AND operator menyalin sedikit ke hasil jika ada di kedua operan.
  ```
  (A & b) akan memberikan 12,
  yaitu 0000 1100
  ```

- ``|``

  biner ATAU operator menyalin sedikit jika ada di salah satu operan.
  ```
  (A | B) menghasilkan 61
  yaitu 0011 1101
  ```

-  ``^``

  operator XOR biner menyalin bit jika diatur dalam satu operan tetapi tidak keduanya
  ```
  (A ^ B) akan memberikan 49
  yaitu 0011 0001
  ```

- ``<<``

  operator shif kiri biner. nilai operan kiri dipindahkan ke kiri dengan jumlah bit yang ditentukan oleh operan kanan.
  ```
  A << 2 akan memberikan 240
  yaitu 1111 0000
  ```

- ``>>``

  operan shift kanan biner. nilai operan kiri dipindahkan ke kanan dengan jumlah bit yang ditentukan oleh operan kanan.
  ```
  A >> 2 akan memberikan 15
  yaitu 0000 1111
  ```

## operator penugasan

| Operator      | deskripsi                             | contoh        |
| :---          | :---:                                 | :---:         |
| ``=``         | operator penugasan sederhana, menetapkan nilai dari operan sisi kanan ke operan sisi kiri     | C = A + B memberikan nilai A + B ke C         |
| ``+=``        | tambahan operator penugasan DAN,ini menambahkan operan kanan ke operan kiri dan menetapkan hasilnya ke operan kiri    | C += A setara dengan C = C + A        |
| ``-=``        | kurangi DAN operator penugasan, ini mengurangi operan kanan dari operan kiri dan menetapkan hasilnya ke operan kiri   | C -= A setara dengan C = C - A        |
| ``*=``        | operator perkalian DAN penugasam, ini mengalikan operan kanan dengan operan kiri dan menetapkan hasilnya ke operan kiri       | c ``*=`` A setara dengan C = C - A |
| ``/=``        | divide AND operator penugasan, ini membagi operan kiri dengan operan kanan dan menetapkan hasilnya ke operan kiri     | C /= a setara dengan C = C / A        |
| ``%=``        | bitwise DAN operator penugasan        | C &= 2 sama dengan C = C & 2  |
| ``^=``        | bitwise ekslusif OR dan operator penugasan    | C ^= 2 sama dengan C = C ^ 2  |P