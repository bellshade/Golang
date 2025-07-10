package searching

// fungsi melakukan pencarian menggunakan algoritam interpolasi searching
// interpolasi searching adalah versi lanjuta dari binary searching yang menentukan
// posisi elemen dengan memperkirakan lokasinya berdasarkan nilai target
//
// Parameter:
// - sortedData  []int: array tipe data int yang sudah terurut ascending
// - tebakan int: nilai yang akan dicari dalam sebuah array
//
// Return:
// - indeks pertama dari nilai yang cocok
// - error jika tidak ditemukan atau kondisi input tidak valid
func pencarianInterpolasi(sortedData []int, tebakan int) (int, error) {
	// jika data kosong, langsung return error karena tidak ada data yang dicari
	if len(sortedData) == 0 {
		return -1, ErrorMessage
	}

	// inisialisasi batas bawah (low) dan atas (high)
	// lowValue = nilai pada index low (elemen pertama)
	// highValue = nilai pada index high (elemen terakhir)
	var (
		low, high           = 0, len(sortedData) - 1
		lowValue, highValue = sortedData[low], sortedData[high]
	)

	// - lowValue != highValue (untuk menghindari terjadinya pembagian dengan nol)
	// - tebakan berada di antara lowValue dan highValue
	// - lowValue <= tebakan && tebakan <= highValue (range valid)
	for lowValue != highValue && (lowValue <= tebakan) && (tebakan <= highValue) {
		// hitung posisi mid secara interpolasi
		// rumus interpolasi
		// mid = low + ((tebakan - lowValue) * (high - low)) / (highValue - lowValue)
		// casting ke float64 digunakan untuk presisi hitungan pecahan
		mid := low + int(float64(float64((tebakan-lowValue)*(high-low))/float64(highValue-lowValue)))

		// jika nilai di mid sama dengan tebakan, kita cek apakah ada duplikasi di
		// sebelumnya
		if sortedData[mid] == tebakan {
			// geser ke kiri sampai ketemu elemen yang bukan tebakan
			// ini nantinya dilakukan untuk menemukan indeks pertama dari nilai
			// yang cocok
			for mid > 0 && sortedData[mid-1] == tebakan {
				mid--
			}
			// setelah ketemu indeks pertama, return dari hasil
			return mid, nil
		}

		// jika nilai di mid lebih besar dari tebakan, geser batas atas
		if sortedData[mid] > tebakan {
			high, highValue = mid-1, sortedData[high]
		} else {
			// jika nilai di mid lebih kecil dari tebakan, geser batas bawah
			low, lowValue = mid+1, sortedData[low]
		}
	}

	// setelah keluar dari looping, kita akan cek sekali lagi apakah
	// lowValue == tebkan, karena bisa saja tebakan berada di tpat di lowValue
	// setlah lowValue setelah iterasi selesai
	if tebakan == lowValue {
		return low, nil
	}
	// jika semua kondisi diatas tidak terpenuhi, maka nilai tidak ditemukan
	return -1, ErrorMessage
}
