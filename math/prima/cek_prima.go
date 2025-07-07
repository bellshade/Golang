package prima

// PembagianTrial memeriksa apakah sebuah bilangan dari n adalah
// bilangan prima
// Parameter:
// 	n - bilangan bulat bertipe integer 64
//
// Return:
//	bool - true jika n adalah bilangan prima, false jika tidak prima
func PembagianTrial(n int64) bool {
	if n < 2 {
		return false
	}
	
	for i := int64(2); i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
