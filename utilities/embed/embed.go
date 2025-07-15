package main

import (
	"embed"
)

// Variabel fileString menyimpan isi dari file single_file.txt sebagai string
// value dari variabel fileString adalah teks/plain dari file tersebut
//
//go:embed folder/single_file.txt
var fileString string

// Variabel fileByte menyimpan isi dari file single_file.txt sebagai []byte (slice of byte)
// Berguna saat kamu membutuhkan data biner seperti untuk operasi hashing, enkripsi, dll.
//
//go:embed folder/single_file.txt
var fileByte []byte

// Variabel folder menyimpan sistem berkas virtual dari beberapa file
// Tipe embed.FS memungkinkan kita membaca banyak file seperti filesystem biasa
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	// mencetak isi dari fileString
	print(fileString)
	// mencetak isi dari fileByte sebagai string
	print(string(fileByte))

	// Membaca file1.hash dari folder
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	// Membaca file2.hash dari folder
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
