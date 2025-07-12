package searching

// melakukan pencarian secar linier (sequential searching)
// untuk menemukan kemunculan pertama dari nilai tertentu (query) di
// dalam array
//
// Parameter:
// - array []int: kumpulan data bertipe integer yang akan dicari
// - query int: nilai yang ingin dicari indeksnya dalam array
//
// Return:
// - int: indeks pertama tempat query ditemukan
// - error: jika query tidak ditemukan, maka throw error
func LinearSearching(array []int, query int) (int, error) {
	// melakukan iterasi terhadap array menggunakan array
	// dimana `i` adalah indeks dan `item` adalah nilai pada posisi tersebut
	for i, item := range array {
		// jika elemen saat ini (item) sama dengan nilai yang dicari
		// maka kembalikan indeksnya dan error nil
		if item == query {
			return i, nil
		}
	}

	// jika semua elemen telah dicek dan tidak ada yang cocok
	// maka kembalikan -1 sebagai indikasi tidak diteumkan, dan throw ErrorMessage
	return -1, ErrorMessage
}
