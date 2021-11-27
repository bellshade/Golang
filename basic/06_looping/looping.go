package main

import "fmt"

func main(){
	var i, j int
	
	// membuat looping pertama
	for i = 2; i < 100; i++{
		// membuat looping kedua
		for j = 2; j <= (i / j); j++ {
			// jika i habis dibagi j, maka bukan bilangan prima
			if (i % j == 0){
				break
			}
		}
		// jika i tidak habis dibagi j, maka i adalah bilangan prima
		if (j > (i / j)){
			fmt.Printf("%d adalah bilangan prima\n", i)
		}
	}
}
