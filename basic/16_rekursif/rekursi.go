package main

import "fmt"

func factorial(i int) int {
	// jika i kurang atau sama dengan 1
	// maka akan mengembalikan nilai 1
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	var i int = 5
	fmt.Printf("faktorial dari %d adalah %d", i, factorial(i))
}
