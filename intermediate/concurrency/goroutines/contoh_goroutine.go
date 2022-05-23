package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	fmt.Println("bellshade")
	time.Sleep(1 * time.Second)
	fmt.Println("selesai")
}

func start() {
	fmt.Println("menjalankan fungsi goroutine")
}
