package searching

import (
	"testing"

	"github.com/bellshade/Golang/algoritma/searching"
)

func TestLinearSearching(t *testing.T) {
	for _, test := range searchTests {
		valueAktual, _ := searching.LinearSearching(test.data, test.key)
		if valueAktual != test.ekspetasi {
			t.Errorf("test `%s` gagal: input array `%v` dengan key `%d`, ekspetasi: `%d`, aktual: `%d`", test.name, test.data, test.key, test.ekspetasi, valueAktual)
		}
	}
}
