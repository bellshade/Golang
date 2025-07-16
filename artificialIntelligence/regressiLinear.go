package artificialIntelligence

import (
	"fmt"
	"math"
)

// DataStatistik adalah stuct yang digunakan untuk menyimpan
// hasil perhitungan statistik dari dua dataset (independent dan dependent)
// anggota meliputi antara lain
// SumX -> jumlah dari seluruh nilai variabel independent
// SumY -> jumlah dari seluruh nilai variabel dependent
// XSquared -> jumlah dari kuadrat variabel independen
// YSquared -> umlah dari kuadrat variabel dependent
// Multiply -> jumlah dari hgasil perkalian antara pasangan nilai dari variabel
type DataStatistik struct {
	SumX     float64
	SumY     float64
	XSquared float64
	YSquared float64
	Multiply float64
}

// menghitung jumlah total dari semua element di dalam slice independenta
// dan dependent
//
// Parameter:
//
//	independent: slice yang berisi data variabel independent
//	dependent: slice yang berisi data variabel dependent
//	stats: pointer ke struct DataStatistik tempat hasil akan disimpan
func Sum(independent, dependent []float64, stats *DataStatistik) {
	var sumX, sumY float64 // inisialisasi penampung jumlah X dan Y

	// looping melalui setiap indeks dari slice independent
	// diasumsikan panjang slice independent dan dependent sama
	for i := range independent {
		sumX += independent[i] // tambahkan nilai X ke sumX
		sumY += dependent[i]   // tambahkan nilai Y ke sumY
	}

	// simpan hasil perhitungan ke struct DataStatistik
	stats.SumX = sumX
	stats.SumY = sumY

	// cetak hasil untuk debug
	fmt.Printf("SumX: %.4f\n", stats.SumX)
	fmt.Printf("SumY: %.4f\n", stats.SumY)
}

// menghitung jumlah dari kuadrat elemen-elemen dalam slice independent dan dependent
//
// Parameter:
//
//	independent -> slice yang berisi variabel independent
//	dependent -> slice yang berisi data variabel dependent
//	stats -> pointer ke struct DataStatistik dimana tempat hasil disimpan
func Squared(independent, dependent []float64, stats *DataStatistik) {
	var xSquared, ySquared float64 // tampung jumlah kuadrat X dan Y

	// looping melalui setiap indeks dari slice independent
	for i := range independent {
		// hitung kuadrat dari nilai X menggunakan math.Pow
		xSquared += math.Pow(independent[i], 2)
		// hitung kuadrat dari nilai Y dengan cara langsung
		ySquared += dependent[i] * dependent[i]
	}

	// simpan hasil perhitungan struct
	stats.XSquared = xSquared
	stats.YSquared = ySquared

	fmt.Printf("hasil X squared: %.4f\n", stats.XSquared)
	fmt.Printf("hasil Y squared: %.4f\n", stats.YSquared)
}

// menghitung jumlah hasil dari perkalian antara setiap parameter pasangan
// lemen dari slice independent dan dependent
//
// Parameter:
//
//	independent -> slice yang berisi data variable independent
//	dependent -> slice yang berisi data variabel dependent
//	stats -> pointe ke struct DataStatistik hasil yang akan disimpan
func Multiply(independent, dependent []float64, stats *DataStatistik) {
	var total float64 // tampung hasil jumlah dari hasil perkalian

	// looping melalui setiap indeks dari slice independent
	for i := range independent {
		// kalikan nilai X[i] dan Y[i], lalu tambahkan nilai total
		total += independent[i] * dependent[i]
	}

	// simpan hasil ke struct
	stats.Multiply = total

	fmt.Printf("hasil kali: %.4f\n", stats.Multiply)
}

// fungsi menghitung model regressi linear sederhana berdasarkan data yang
// telah dihitung sebelumnya, model regressi berbentuk: y = a + bx
// di mana:
// b adalah slope (kemiringan garis)
// a adalah intercept (titik potong garis dengan sumbu Y)
//
// Parameter:
//   - independent -> slice yang berisi data variabel independent
//   - dependent -> slice yang bersisi data variabel dependent
//   - stats -> instance DataStatistik
func RegressiLinear(independent, dependent []float64, stats DataStatistik) {
	// jumlah pasangan data, n
	n := float64(len(independent))

	// hitung jumlah kuadrat dari deviasi XY (sumSquareXY)
	// rumus = sum(xy) - (sum(x) * sum(y)) / n
	sumSquareXY := stats.Multiply - (stats.SumX*stats.SumY)/n

	// hitung jumlah kuadrat dari deviasi XX (sumSquareXX)
	// rumus = sum(x^2) - (sum(x)) ^ 2 / n
	sumSquareXX := stats.XSquared - (math.Pow(stats.SumX, 2))/n

	fmt.Printf("hasil sumSquareXY: %.4f\n", sumSquareXY)
	fmt.Printf("hasil sumSquareXX: %.4f\n", sumSquareXX)

	// hitung slope (b) menggunakan rumus = sumSquareXY / sumSquareXX
	slope := sumSquareXY / sumSquareXX
	// hitung intercept (a) menggunakan rumus: (sum(y) - b * sum(x)) / n
	intercept := (stats.SumY - (slope * stats.SumX)) / n

	fmt.Printf("model regresi linear: y = %.4f + %.4fx\n", intercept, slope)
}
