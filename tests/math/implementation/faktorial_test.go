package math_implementation

import (
	"fmt"
	"testing"

	"github.com/bellshade/Golang/math_implementation/faktorial"
)

type fungsiFaktorial func(int) (int, error)

var implementasi = map[string]fungsiFaktorial{
	"Iterasi":  faktorial.Iterasi,
	"Rekursif": faktorial.Rekursif,
	"Tree":     faktorial.TreeBin,
}

var testCaseFactorial = []struct {
	n         int
	ekspetasi int
}{
	{0, 1},
	{1, 1},
	{6, 720},
	{9, 362880},
}

func TestFaktorial(t *testing.T) {
	for implName, implFunction := range implementasi {
		t.Run(implName+" error input negatif", func(t *testing.T) {
			_, error := implFunction(-1)
			if error != faktorial.ArgumentNegatif {
				t.Errorf("tidak ada error dari input negatif")
			}
		})

		for _, tc := range testCaseFactorial {
			t.Run(fmt.Sprintf("%s dengan input %d", implName, tc.n), func(t *testing.T) {
				aktual, err := implFunction(tc.n)
				if err != nil {
					t.Errorf("add error")
				}
				if aktual != tc.ekspetasi {
					t.Errorf("ekspetasi: %d, aktual: %d", tc.ekspetasi, aktual)
				}
			})
		}
	}
}
