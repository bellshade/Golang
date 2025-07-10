package searching

import (
	"fmt"
	"math"
)

// ternaryMaxSearching mencari nilai maks lokal dari fungsi f(x) dalam interval [a, b]
// dengan toleransi epsilon menggunakan ternary searching
//
// algoritma
// - interval [a, b] dibagi menjadi tiga bagian yang sama panjang
// - evaluasi nilai f(x) pada dua titik tengah: kiri dan kanan
// - jika f(kiri) < f(kanan), maka nilai maks berada dis sisi kanan
// proses ini akan berlanjut secara rekursif hingga interval lebih kecil dari epsilon
//
// Parameter:
// a - batas kiri interval (float64)
// b - batas kanan interval (float64)
// epsilon - toleransi akurasi hasil (float64)
// f - fungsi matematis satu variable bertipe func(x float64) float64
//
// Return:
// float64 - nilai maks dari f(x) dalam unterval [a, b]
// error - error jika input tidak valid atau terjadi masalah selama eksekusi
func ternaryMaxSearching(a, b, epsilon float64, f func(x float64) float64) (float64, error) {
	if a == math.Inf(-1) || b == math.Inf(1) {
		return -1, fmt.Errorf("interval harus memiliki angka float64")
	}
	if math.Abs(a-b) <= epsilon {
		return f((a + b) / 2), nil
	}

	kiri := (2*a + b) / 3
	kanan := (a + 2*b) / 3
	if f(kiri) < f(kanan) {
		return ternaryMaxSearching(kiri, b, epsilon, f)
	}

	return ternaryMaxSearching(a, kanan, epsilon, f)
}

// ternaryMinSearching mencari nilai minimum lokal dari fungsi(x) dalam interval [a, b]
//
// algoritma:
// - interval [a, b] dibagi menjadi tiga bagian sama panjang
// - evaluasi nilai f(x) pada dua titik tengah: kiri dan kanan
// - jika f(kiri) > f(kanan), maka nilai minimum berbeda di sisi kanan
// proses berlanjut secara rekursif hingga interval lebih kecil dari epsilon
func ternaryMinSearching(a, b, epsilon float64, f func(x float64) float64) (float64, error) {
	if a == math.Inf(-1) || b == math.Inf(1) {
		return -1, fmt.Errorf("interval harus memiliki angka float64")
	}

	if math.Abs(a-b) <= epsilon {
		return f((a + b) / 2), nil
	}

	kiri := (2*a + b) / 3
	kanan := (a + 2*b) / 3
	if f(kiri) > f(kanan) {
		return ternaryMinSearching(kiri, b, epsilon, f)
	}
	return ternaryMinSearching(a, kanan, epsilon, f)
}
