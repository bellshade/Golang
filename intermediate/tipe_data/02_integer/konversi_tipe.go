package main

import "fmt"

func main() {
	var x int32
	var y uint32 // 0 sampai 4294967295
	var z uint8  // 0 sampai 255

	x = 26700

	y = uint32(x) // data ditampilkan secara penuh karena size terpenuhi

	z = uint8(x) // data hilang karena size sudah melewati batas

	fmt.Println(y, z)
}
