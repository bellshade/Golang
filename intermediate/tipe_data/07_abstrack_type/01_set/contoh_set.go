// Package main
package main

import (
	"fmt"
)

// Definisi generic Set
type Set[K comparable] map[K]struct{}

// Tambah elemen
func (s Set[K]) Add(val K) {
	s[val] = struct{}{}
}

// Hapus elemen
func (s Set[K]) Remove(val K) {
	delete(s, val)
}

// Cek apakah elemen ada
func (s Set[K]) Contains(val K) bool {
	_, ok := s[val]
	return ok
}

// Tambah semua elemen dari set lain
func (s Set[K]) AddAll(other Set[K]) {
	for k := range other {
		s.Add(k)
	}
}

// Pertahankan hanya elemen yang juga ada di set lain
func (s *Set[K]) RetainAll(other Set[K]) {
	sNew := Set[K]{}
	for v := range other {
		if s.Contains(v) {
			sNew.Add(v)
		}
	}
	*s = sNew
}

// Hapus semua elemen yang ada di set lain
func (s Set[K]) RemoveAll(other Set[K]) {
	for v := range other {
		s.Remove(v)
	}
}

// Hitung jumlah elemen
func (s Set[K]) Len() int {
	return len(s)
}

// Contoh penggunaan
func main() {
	// 1. Membuat Set dan menambahkan elemen
	set := Set[string]{}
	set.Add("A")
	set.Add("B")
	set.Add("C")
	fmt.Println("Set awal:", set)

	// 2. Mengecek apakah elemen ada
	fmt.Println("Apakah ada 'A'?", set.Contains("A"))
	fmt.Println("Apakah ada 'Z'?", set.Contains("Z"))

	// 3. Menghapus elemen
	set.Remove("A")
	fmt.Println("Set setelah menghapus 'A':", set)

	// 4. Union (gabung set lain)
	set2 := Set[string]{}
	set2.Add("B")
	set2.Add("X")
	set2.Add("Y")
	set.AddAll(set2)
	fmt.Println("Set setelah AddAll set2:", set)
}
