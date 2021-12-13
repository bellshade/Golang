package main

import "fmt"

type Buku struct {
	judul   string
	penulis string
	id_buku int
}

func main() {
	var Buku1 Buku

	Buku1.judul = "Harry potter the dathly hallows"
	Buku1.penulis = "J.K Rowling"
	Buku1.id_buku = 1761

	fmt.Printf("judul buku %s\n", Buku1.judul)
	fmt.Printf("dengan penulis %s\n", Buku1.penulis)
	fmt.Printf("id buku adalah %d\n", Buku1.id_buku)
}
