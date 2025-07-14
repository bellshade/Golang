package kompresi

import (
	"sort"
	"testing"

	"github.com/bellshade/Golang/kompresi"
)

func FrekuensiSimbol(pesan string) []kompresi.FrekuensiSimbol {
	hitungRune := make(map[rune]int)
	for _, s := range pesan {
		hitungRune[s]++
	}
	listFrekuensi := make([]kompresi.FrekuensiSimbol, len(hitungRune))
	i := 0
	for s, n := range hitungRune {
		listFrekuensi[i] = kompresi.FrekuensiSimbol{Simbol: s, Frekuensi: n}
		i++
	}
	sort.Slice(listFrekuensi, func(i, j int) bool {
		return listFrekuensi[i].Frekuensi < listFrekuensi[j].Frekuensi
	})
	return listFrekuensi
}

func TestFungsiHuffman(t *testing.T) {
	pesan := []string{
		"bellshade bahasa pemograman Golang",
		"indonesia adalah negara kesatuan berbentuk republik",
	}

	for _, kata := range pesan {
		t.Run("huffman: "+kata, func(t *testing.T) {
			tree, _ := kompresi.TreeHuffman(FrekuensiSimbol(kata))
			codes := make(map[rune][]bool)
			kompresi.HuffmanEncoding(tree, nil, codes)
			messageTerkode := kompresi.HuffmanEncode(codes, kata)
			pesanHuffmanDecoded := kompresi.HuffmanDecode(tree, tree, messageTerkode, "")
			if pesanHuffmanDecoded != kata {
				t.Errorf("error: %q\n ekspetasi: %q", pesanHuffmanDecoded, kata)
			}
		})
	}
}
