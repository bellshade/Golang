package main

import "fmt"

func main(){
	// membuat variabel
    var a int = 10
    var b int = 20
    var hasil int
    
	// memanggil fungsi
    hasil = hitung(a, b)

	// menampilkan hasil
    fmt.Printf("hasilnya adalah: %d\n", hasil)
    
}

// membuat fungsi
func hitung(angka1, angka2 int) int{
	// membuat variabel
    var hasil int 
	// menghitung antara 2 variabel
    hasil = angka1 + angka2

	// mengembalikan nilai
    return hasil
}
