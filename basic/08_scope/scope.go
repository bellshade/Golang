package main

import "fmt"

// membuat variabel global
var a int = 20

func main() {
	// membuat variabel lokal
	// variabel lokal ini tidak
	// bisa diakses dari luar fungsi
	var a int = 10
	var b int = 20

	// membuat variabel yang memanggil
	// fungsi tambah
	var c int = tambah(a, b)

	// menampilkan hasil
	fmt.Printf("hasil %d", c)
}

func tambah(a, b int) int {
	// membuat pengembalian nilai
	// berupa hasil pertambahan
	// antara kedua variabel
	return a + b
}
