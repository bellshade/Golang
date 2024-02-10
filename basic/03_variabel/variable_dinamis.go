package main

import "fmt"

func main() {
	// membuat variabel dengan tipe data
	// float64
	var angka_saya float64 = 29.2

	// deklarasi tipe data dengan operator
	// walrus
	angka_lain := 12

	fmt.Println(angka_saya)
	fmt.Println(angka_lain)

	// melihat tipe data dari
	// kedua variabel
	fmt.Printf("tipe data dari angka_saya adalah %T\n", angka_saya)
	fmt.Printf("tipe data dari angka_lain adalah %T\n", angka_lain)
}
