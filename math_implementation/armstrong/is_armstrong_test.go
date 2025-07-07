package armstrong

import "testing"

var testCase = []struct {
	name     string
	input    int
	expected bool
}{
	{
		"angka negatif: bukan angka armstrong",
		-140,
		false,
	},
	{
		"angka positif: bukan angka armstrong",
		23,
		false,
	},
	{
		"angka armstrong paling kecil",
		0,
		true,
	},
	{
		"bilangan armstring terkecil dengan lebih dari satu digit",
		153,
		true,
	},
	{
		"angka armstrong acak",
		407,
		true,
	},
}

func TestArmstrong(t *testing.T) {
	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			funcResult := cekArmstrong(test.input)
			if test.expected != funcResult {
				t.Errorf("expect jawaban '%t' untuk nomor '%d' tapi jawaban yang diberikan %t", test.expected, test.input, funcResult)
			}
		})
	}
}
