package main

import "fmt"

func main() {
	/**
	 * Percabangan dengan if else Dasar
	 */

	// Deklarasi Variable
	numChild := 5

	if numChild > 2 { // numChild > 2 merupakan condition statement
		/**
		 * Baris ini hanya akan di eksekusi ketika condition statement terpenuhi
		 * atau menghasilkan nilai true
		 */
		fmt.Println("Banyak anak banyak rejeki")
	} else {
		/**
		 * Baris ini dijalankan ketika condition statement tidak terpenuhi
		 * atau menghasilkan nilai false
		 */
		fmt.Println("Dua anak cukup")
	}

	/**
	 * Percabangan dengan Temporary Variable
	 */
	if numChild2 := 5; numChild2 > 3 {
		// numChild2 hanya dapat digunakan pada blok kode if else, diluar itu maka akan error
		fmt.Println(numChild2)
		fmt.Println("Lebih dari 3")
	} else {
		fmt.Println(numChild2)
		fmt.Println("Kurang dari 3")
	}

	/**
	 * Ketika kita mencoba mengakses variabel numChild2, maka akan muncul error undeclared name
	 * Silakan uncomment baris 43 untuk mengetahui error tersebut
	 */
	// fmt.Println(numChild2)

	/**
	 * Percabangan dengan Switch case
	 */

	// Switch case dasar
	switch numChild { // numChild ada sebuah nilai yang akan dicocokkan dengan case-case yang ada di dalam switch statement
	case 5:
		fmt.Println("Switch case 5") // jika numChild bernilai 5, maka akan mengeksekusi kode tersebut
	case 3:
		fmt.Println("Switch case 3") // jika numChild bernilai 3, maka akan mengeksekusi kode tersebut
	default:
		fmt.Println("Switch case default") // jika numChild bukan bernilai 5 dan 3, maka akan mengeksekusi kode tersebut
	}

	// Switch case seperti if else
	// Agar dapat menggunakan pengecekan kondisi seperti di percabangan if else, maka tidak perlu memberikan variabel atau nilai setelah statement `switch`
	// Jika setelah statement `switch` tidak ada nilai yang akan dicek, maka kita dapat menggunakan pengecekan kondisi seperti di bawah ini pada setiap statement `case`
	switch {
	case numChild == 5:
		fmt.Println("Switch case sama dengan 5")
	case numChild < 5:
		fmt.Println("Switch case kurang dari 5")
	default:
		fmt.Println("Switch case lebih dari 5")
	}
}
