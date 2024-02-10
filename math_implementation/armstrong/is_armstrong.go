// angka armstrong
// memeriksa apakah nomor asyang diberikan adalah nomor armstrong
// atau bukan
// penjelasan tentang angka armstrong
// https://medium.com/@achmadg6/apa-itu-armstrong-number-c60af4693ae

package armstrong

import (
	"math"
	"strconv"
)

func cekArmstrong(angka int) bool {
	var digit_kanan int
	var sum int = 0
	var tempNum int = angka

	// untuk mendapatkan jumlah digit
	// di dalam nomor
	length := float64(len(strconv.Itoa(angka)))

	// dapatkan digit paling kanan dan putuskan loop
	// setelah semua digit diulang

	for tempNum > 0 {
		digit_kanan = tempNum % 10
		sum += int(math.Pow(float64(digit_kanan), length))

		// perbarui digit input dikurangi yang diproses
		// paling kanan
		tempNum /= 10
	}

	return angka == sum
}
