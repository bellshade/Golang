package logaritma

func LogBase2(n uint32) uint32 {
	base := [5]uint32{0x2, 0xC, 0xF0, 0xFF00, 0xFFFF0000}
	eksponen := [5]uint32{1, 2, 4, 8, 16}
	var hasil uint32
	for i := 4; i >= 0; i-- {
		if n&base[i] != 0 {
			n >>= eksponen[i]
			hasil |= eksponen[i]
		}
	}
	return hasil
}
