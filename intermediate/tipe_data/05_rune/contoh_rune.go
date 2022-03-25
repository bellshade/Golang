package main

import (
	"fmt"
	"reflect"
)

func main() {
	// membuat rune
	runePertama := 'B'
	runeKedua := 'g'

	// menampilkan contoh rune yang sudah dibuat
	fmt.Printf("contoh rune pertama: %c; unicode: %U; tipe: %s", runePertama, runePertama, reflect.TypeOf(runePertama))

	fmt.Printf("contoh rune pertama: %c; unicode: %U; tipe: %s", runeKedua, runeKedua, reflect.TypeOf(runeKedua))
}
