package main

import (
	"fmt"
	"sync"
	"time"
)

// jika WaitGroup digunakan sebagai parameter, pastikan parameter tersebut adalah pointer
func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done() // Decreasing the count in the WaitGroup after the goroutine is finished

	fmt.Printf("Worker %d starting \n", x)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done \n", x)
}

func main() {
	// Initialize WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1) // Adding 1 to the count of goroutines to be waited
		go worker(i, &wg)
	}

	wg.Wait() // Waiting until all goroutines are finished

	fmt.Println("All worker are finished")
}
