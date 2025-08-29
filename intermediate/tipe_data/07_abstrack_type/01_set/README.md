# Set di Golang (Generic Abstract Type)

Dalam bahasa Go, **`Set`** bukanlah tipe bawaan. Namun, kita bisa membuat **abstraksi tipe `Set`** sendiri menggunakan **`map[K]struct{}`**.
Tipe ini sangat berguna untuk **menyimpan elemen unik** (tidak duplikat), melakukan operasi seperti **pengecekan keanggotaan**, **union**, **intersection**, dan **difference**.

---

## **Deklarasi Tipe Set**

```go
type Set[K comparable] map[K]struct{}
```

- `K comparable` → artinya kunci (`key`) harus bertipe yang bisa dibandingkan (`string`, `int`, dll).
- `struct{}` dipakai sebagai nilai karena ringan (tanpa konsumsi memori tambahan).

---

## **Operasi Dasar pada Set**

### 1. Membuat Set & Menambahkan Elemen

```go
set := Set[string]{}
set.Add("A")
set.Add("B")
```

### 2. Mengecek Keanggotaan

```go
fmt.Println(set.Contains("A")) // true
fmt.Println(set.Contains("C")) // false
```

### 3. Menghapus Elemen

```go
set.Remove("A")
fmt.Println(set.Contains("A")) // false
```

### 4. Menambahkan Semua Elemen dari Set Lain (Union)

```go
set2 := Set[string]{}
set2.Add("X")
set2.Add("Y")

set.AddAll(set2)
fmt.Println(set) // {"B","X","Y"}
```

### 5. Menyisakan Elemen yang Sama (Intersection)

```go
set.RetainAll(set2)
fmt.Println(set) // {"X","Y"}
```

### 6. Menghapus Semua Elemen yang Ada di Set Lain (Difference)

```go
set3 := Set[string]{}
set3.Add("X")

set.RemoveAll(set3)
fmt.Println(set) // {"Y"}
```

---

## **Keterbatasan**

- Tipe generik `Set[K]` hanya bisa digunakan untuk **tipe yang comparable**.
- Tidak bisa langsung menghapus `Set[int]` dari `Set[string]`, karena beda tipe.

---

## **Perbedaan dengan Slice**

| Aspek             | `Set`                           | `Slice`               |
| ----------------- | ------------------------------- | --------------------- |
| Duplikat Elemen   | Tidak boleh                     | Boleh                 |
| Operasi Pencarian | Sangat cepat (O(1))             | Lebih lambat (O(n))   |
| Urutan Elemen     | Tidak terjamin                  | Terjaga sesuai urutan |
| Cocok untuk       | Membership, union, intersection | Data berurutan, list  |
