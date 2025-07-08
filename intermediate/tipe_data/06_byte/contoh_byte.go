package main

import (
	"fmt"
	"reflect"
)

func main() {
	// membuat byte
	var b byte
	b = 200
	fmt.Printf("byte: %c; nilai: %d; tipe: %s\n", b, b, reflect.TypeOf(b))
	var b2 byte
	b2 = 'C'
	fmt.Printf("byte: %c; nilai: %d; tipe: %s\n", b2, b2, reflect.TypeOf(b2))

	// membuat bytes (sekumpulan data byte)
	var bs []byte
	bs = []byte{72, 101, 108, 108, 111}
	fmt.Println("bytes", reflect.TypeOf(bs))
	fmt.Println("bytes menjadi string ->", string(bs))
}
