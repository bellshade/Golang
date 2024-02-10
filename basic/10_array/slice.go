package main

import "fmt"

func main() {
	// membuat array

	var makanan = []string{"ayam", "ikan", "mie", "telur", "bakso"}

	// membut slice dan kemudian memasukkan dalam varibel
	// terbaru
	// mengambil dari index 0 ke 3
	// ayam ikan mie dan telur

	var makananSaya = makanan[0:3]

	// print hasil slice

	fmt.Println(makananSaya)
}
