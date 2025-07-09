package searching

import "testing"

func TestLinearSearching(t *testing.T) {
	for _, test := range searchTests {
		valueAktual, _ := linearSearching(test.data, test.key)
		if valueAktual != test.ekspetasi {
			t.Errorf("test `%s` gagal: input array `%v` dengan key `%d`, ekspetasi: `%d`, aktual: `%d`", test.name, test.data, test.key, test.ekspetasi, valueAktual)
		}
	}
}
