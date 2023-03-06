# Generic

Generic adalah kemampuan menambahkan parameter type saat membuat function
Berbeda dengan tipe data yang biasa kita gunakan di function, generic memungkinkan kita bisa mengubah-ubah bentuk tipe data sesuai dengan yang kita mau.

### Type Inferance

Type Inferance merupakan fitur yang dimana kita tidak perlu lagi menyebutkan Type Parameter ketika memanggil kode generic
Tipe data type Parameter bisa dibaca secara otomatis misal dari parameter yang kita kirim
Contoh nya adalah
Before:

```go
func Length[T any](param T) T {
	return param
}

func main() {
  result := Length[string]("Hello")
	fmt.Println(result)
}
```

After

```go
func Length[T any](param T) T {
	return param
}

func main() {
  result := Length("Hello")
	fmt.Println(result)
}
```

Bisa dilihat di fungsi main before afternya, kita tidak perlu lagi menyebutkan generic type didalam kurung kotak tersebut karena golang sudah tau data yang dikirimkan di parameter dan secara otomatis golang tau itu adalah tipe data string
