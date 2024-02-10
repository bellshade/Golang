package main

import "fmt"

func main() {
	// membuat 2 variabel yang bertipe data integer
	var angkaPertama int = 12
	var angkaKedua int = 3

	// membuat tipe data float sebagai
	// variabel hasil dari konversi tipe daata
	// antara angkaPertama dan angkaKedua
	var hasil float32

	hasil = float32(angkaPertama) / float32(angkaKedua)

	fmt.Printf("hasil pembagian adalah %f\n", hasil)
}
