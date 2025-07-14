package kompresi

import (
	"fmt"
)

// node merepresentasikan simpla dalam pohon huffman
// setiap node memiliki
// - pointer ke child kiri
// - pointer ke child kanan
// - karakter ASCII jika ini nantinya adalah leaf node
// - frekuensi kemunculan dari simbol
type Node struct {
	kiri   *Node
	kanan  *Node
	simbol rune
	bobot  int
}

// FrekuensiSimbol menyimpan informasi kemunculan suatu simbol
type FrekuensiSimbol struct {
	Simbol    rune // karakter ASCII
	Frekuensi int  // jumlah kemunculan simbol tersebut
}

// fungsi membangun pohon huffman berdasarkan daftar list frekuensi simbol
// menggunakan dua antrian untuk efisiensi tanpa heap
//
// Parameter:
//
//	listFrekuensi []FrekuensiSimbol daftar list pasangan simbol
//
// Return:
// *Node: pointer ke root pohon huffman
// error: jika input tidak valid
func TreeHuffman(listFrekuensi []FrekuensiSimbol) (*Node, error) {
	if len(listFrekuensi) < 1 {
		return nil, fmt.Errorf("TreeHuffman: list tidak boleh kosong")
	}

	// inisialisasi antrian pertama (q1) dengan semua leaf node
	q1 := make([]Node, len(listFrekuensi))
	// q2 digunakan untuk node internal
	q2 := make([]Node, 0, len(listFrekuensi))

	for i, x := range listFrekuensi {
		q1[i] = Node{
			kiri: nil, kanan: nil, simbol: x.Simbol, bobot: x.Frekuensi,
		}
	}

	// bangun pohon hingga tersisa satu node
	for len(q1)+len(q2) > 1 {
		var node1, node2 Node
		// ambil node dengan bobot nilai terkecil
		node1, q1, q2 = least(q1, q2)
		node2, q1, q2 = least(q1, q2)
		// buat node internal baru sebagai parent dari kedua node
		node := Node{kiri: &node1, kanan: &node2, simbol: -1, bobot: node1.bobot + node2.bobot}

		// tambahin nde ke antrian kedua
		q2 = append(q2, node)
	}

	// return root pohon huffman
	if len(q1) == 1 {
		return &q1[0], nil
	}
	return &q2[0], nil
}

// fungsi mengambil node dengan bobot nilai terkecil dari antara dua antrian
//
// Parameter:
//
//	q1 []Node: Antrian awal dengan leaf node
//	q2 []Node: Antrian awal untuk node internal
//
// Return:
// Node: node ndegan bobot terkecil
// []Node: sisa antrian q1 setelah ekstrasi
// []Node: sisa antrian q2 setelah ekstrasi
func least(q1 []Node, q2 []Node) (Node, []Node, []Node) {
	if len(q1) == 0 {
		// ambil dari q2
		return q2[0], q1, q2[1:]
	}
	if len(q2) == 0 {
		// ambil dari q1
		return q1[0], q1[1:], q2
	}
	if q1[0].bobot <= q2[0].bobot {
		// ambil dari q1 karena lebih kecil
		return q1[0], q1[1:], q2
	}
	// ambil dari q2 karena lebih kecil
	return q2[0], q1, q2[1:]
}

// fungsi yang melakukan traversal pohon huffman untuk membuat tabel
// kode biner
//
// Parameter:
//
//	node *Node: root dari pohon huffman
//	prefix []bool: jalur dari root ke leaf (false = kiri, true = kanan)
//	codes map[rune][]bool: peta simbol ke representasi bitnya
func HuffmanEncoding(node *Node, prefix []bool, codes map[rune][]bool) {
	// jika node adalah leaf (memiliki simbol), simpan kode biner
	if node.simbol != -1 {
		codes[node.simbol] = prefix
		return
	}

	// rekursi ke child kiri
	prefixKiri := make([]bool, len(prefix))
	copy(prefixKiri, prefix)
	prefixKiri = append(prefixKiri, false) // false = maka ke kiri
	HuffmanEncoding(node.kiri, prefixKiri, codes)

	// rekursi ke child kanan
	prefixKanan := make([]bool, len(prefix))
	copy(prefixKanan, prefix)
	prefixKanan = append(prefixKanan, true) // true = maka ke kanan
	HuffmanEncoding(node.kanan, prefixKanan, codes)
}

// fungsi mengubah string menjadi rangkaian bit berdasarkan kode huffman
//
// Parameter:
//
//	codes map[rune][]bool: tabel pemetaan simbol ke kode biner
//	in string: input teks yang akan dikodekan
//
// Return:
//
//	[]bool: representasi biner dari teks input
func HuffmanEncode(codes map[rune][]bool, in string) []bool {
	out := make([]bool, 0)
	for _, s := range in {
		out = append(out, codes[s]...)
	}
	return out
}

// fungsi mengubah rangkaian bit kembali ke bentuk asli menggunakan pohon huffman
//
// Parameter:
//
//	root *Node: root poohn huffman
//	current *Node: node saat ini dalam traversal
//	in []bool: input biner yang akan didekode
//	out string: output hasil dekoding
//
// Return:
//
//	string: hasil dekoding dari bit input
func HuffmanDecode(root, current *Node, in []bool, out string) string {
	// jika sudah mencapai leaf node, tambahkan simbol ke hasil
	if current.simbol != -1 {
		out += string(current.simbol)
		// lanjutkan dari akar lagi untuk karakter berikutnya
		return HuffmanDecode(root, root, in, out)
	}

	// jika input habis, return hasil
	if len(in) == 0 {
		return out
	}

	// geser ke child kiri atau kanan sesuai bit saat ini
	if in[0] {
		return HuffmanDecode(root, current.kanan, in[1:], out)
	}
	return HuffmanDecode(root, current.kiri, in[1:], out)
}
