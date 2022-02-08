package main

import "fmt"

func main() {
	// contoh program boolean sederhana
	var a bool = true
	var b bool = false

	x := 4
	y := 20

	nilai := 90

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// contoh pada tipe data angka integer
	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x <= y:", x <= y)
	fmt.Println("x >= y:", x >= y)

	// menggunakan pernyataan majemuk
	fmt.Println(!((-0.2 > 1.4) && ((0.8 < 3.1) || (0.1 == 0.1))))

	// menggunakan operator logika
	if nilai > 80 {
		fmt.Println("lewat ujian")
	} else {
		fmt.Println("tidak lewat ujian")
	}

}
