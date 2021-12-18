package main

import "fmt"

func main() {
	var angka []int
	printSlice(angka)

	/* append nile slice */
	angka = append(angka, 0)
	printSlice(angka)

	/* menambahkan 1 elemen ke dalam slice */
	angka = append(angka, 1)
	printSlice(angka)

	/* menambahkan lebih dari satu elemen */
	angka = append(angka, 2, 3, 4)
	printSlice(angka)

	/* membuat slice angka1 dengan kapasitas ganda */
	angka1 := make([]int, len(angka), (cap(angka))*2)

	/* menyalin konten dari angka ke angka1 */
	copy(angka1, angka)
}

func printSlice(x []int) {
	fmt.Printf("len%d cap=%d slice=%v\n", len(x), cap(x), x)
}
