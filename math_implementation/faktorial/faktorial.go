package faktorial

import "errors"

var ArgumentNegatif = errors.New("argumen input harus berupa bilangan bulat non-negatif")

// fungsi untuk menghitung faktorial secara bruteforce
func Iterasi(n int) (int, error) {
	if n < 0 {
		return 0, ArgumentNegatif
	}

	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil, nil
}

// fungsi menghitung faktorial secara rekursif
func Rekursif(n int) (int, error) {
	if n < 0 {
		return 0, ArgumentNegatif
	}

	if n <= 1 {
		return 1, nil
	}

	prev, _ := Rekursif(n - 1)
	return n * prev, nil
}

// menghitung faktorial menggunakan tree biner (divide and conquer)
func TreeBin(n int) (int, error) {
	if n < 0 {
		return 0, ArgumentNegatif
	}

	if n == 0 {
		return 1, nil
	}

	if n == 1 || n == 2 {
		return n, nil
	}

	return produkTree(2, n), nil
}

func produkTree(l int, r int) int {
	if l > r {
		return 1
	}

	if l == r {
		return l
	}

	if r-l == 1 {
		return l * r
	}

	m := (l + r) / 2
	return produkTree(l, m) * produkTree(m+1, r)
}
