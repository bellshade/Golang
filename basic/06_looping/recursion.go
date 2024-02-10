package main

import "fmt"

func main() {
	var n int = 3
	fmt.Printf("%d! = %d", n, factorial_number(n))
}

// fungsi factorial_number akan mencari hasil faktorial dari
// sebuah bilangan dengan cara perulangan recursive
func factorial_number(n int) int {
	// print stack (perulangan) fungsi
	fmt.Printf("bilangan faktorial %d\n", n)

	// jika n atau parameter sama dengan 0
	// tidak ada perulangan lagi (fungsi telah selesai dijalankan)
	if n == 0 {
		return 1
	}

	// memanggil kembali fungsi factorial_number
	// dikarenakan nilai n belum sama dengan 0
	return n * factorial_number(n-1)
}
