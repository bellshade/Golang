package main

import "fmt"

// pada materi ini kita akan belajar tentang variable pada golang
// variable adalah tempat untuk menyimpan data.
// golang merupakan bahasa pemrograman static type yang artinya satu variable hanya dapat menampung
// satu data dengan tipe data tertentu
// tapi golang juga merupakan bahasa pemrograman yang dynamic type, yang mana bisa menyimpan nilai ke variable
// tanpa harus mendeklarasikan tipe data dari nilai tersebut

func main() {
//	pada golang terdapat beberapa cara untuk membuat variable.

//	1 (static type)
	var nama string = "Manusia"
	fmt.Println(nama)

//	2 (shorthand) -> cara membuat variable dengan cara yang lebih singkat
	alamat := "Bumi"
	fmt.Println(alamat)

//	3 (static type)
	// zeros value, adalah nilai default dari tipe data yang digunakan dalam golang,
	// nilai ini akan digunakan jika variable tidak diberikan nilai awal, contoh
	// int -> 0
	// float -> 0.0
	// bool -> false
	// string -> ""
	var jenisKelamin string
	fmt.Println(jenisKelamin) // ini akan berisi string kosong
	jenisKelamin = "laki - laki"
	fmt.Println(jenisKelamin) // ini akan berisi laki - laki

//	4 (static type)
/*
	cara ini dilakukan jika ada beberapa variable dengan tipe data yang sama, tapi nilai variable tersebut tidak
	ditentukan diawal
*/
	var menikah,dewasa,tidur bool
	menikah = false
	dewasa = true
	tidur = false
	fmt.Println(menikah, dewasa, tidur)
}