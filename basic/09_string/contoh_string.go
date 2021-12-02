package main

import "fmt"

func main() {
	var pesan = "bellshade"

	fmt.Printf("normal string ")
	fmt.Printf("%s", pesan)
	fmt.Printf("\n")
	fmt.Printf("hex byte: ")

	for i := 0; i < len(pesan); i++ {
		fmt.Printf("%x", pesan[i])
	}

	fmt.Printf("\n")
	const contohPesan = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Printf("string menggunakan quote ")
	fmt.Printf("%+q", contohPesan)
	fmt.Printf("\n")
}
