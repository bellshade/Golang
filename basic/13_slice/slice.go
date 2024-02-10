package main

import "fmt"

func main() {
	/* membuat slice */
	angka := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(angka)

	/* print original slice */
	fmt.Println("angka ==", angka)

	/* print dari subslice yng dimulai dari index 1 */
	fmt.Println("angka[1:4] ==", angka[1:4])

	fmt.Println("angka[4:] ==", angka[4:])

	angka1 := make([]int, 0, 5)
	printSlice(angka1)

	angka2 := angka[2:]
	printSlice(angka2)

	angka3 := angka[2:5]
	printSlice(angka3)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
