package matriks

import "fmt"

// tipe alias untuk representasi matriks dua dimensi
// menggunakan slice of slice dari int, [][]int
//
// informasi:
// setiap elemen luar ([]int) represenatsi dari satu baris dalam matriks
// setiap elemen dalam (int) nilai skalar pada kolom terentu di baris
type Matriks [][]int

// fungsi melakukan pertambahan dari dua buah matriks (a dan b)
// sesuai aturan matematika: C = A + B, dimana c[i][j] = a[i][j] + b[i][j]
//
// Parameter:
//  - a: matriks pertama
//  - b: matriks kedua
//
// Return:
//  - matriks hasil dari penjumlahan
//  - error jika dimensi matriks tidak kompatibel
//
// - function bersifat immutable: tidak mengubah input a dan b
// - menggunakan deep copy untuk hasil tidak berbagi referensi memori dengan input
func Tambah(a, b Matriks) (Matriks, error) {
	// sebelum melanjutkan pastikan kedua matriks memiliki jumlah baris
	// yang sama
	if len(a) != len(b) {
		return nil, fmt.Errorf("jumlah baris tidak sama")
	}

	// buat slice baru dengan panjang yang sama seperti baris di matriks a
	// ini hanya membuat slice diluar (baris)
	hasil := make(Matriks, len(a))

	// iterasi setiap baris matriks a menggunakan loop range
	for i, barisA := range a {
		// ambil baris yang sesuai dari matriks b
		barisB := b[i]
		
		// validasi jumlah kolom pada baris ke-i
		if len(barisA) != len(barisB) {
			return nil, fmt.Errorf("jumlah kolom pada baris %d tidak sama", i)
		}

		// siapkan slice baru untuk menyimpan hasil penjumlahan baris ke-i
		// panjangnya disesuaikan dengan jumlah kolom
		hasilBaris := make([]int, len(barisA))

		// iterasi setiap elemen dari baris menggunakan range
		for j, valueA := range barisA {
			// jumlahkan elemen yang bersesuaian dari matriks a dan b
			hasilBaris[j] = valueA + barisB[j]
		}
		hasil[i] = hasilBaris
	}
	return hasil, nil
}
