package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var total1 int64

	// 1. Tanpa sinkronisasi, berpotensi race condition
	for range 10000 {
		go func() {
			total1++
		}()
	}

	// 2. Dengan Mutex - penguncian eksplisit
	var total2 int64
	mx := sync.Mutex{}
	for range 10000 {
		go func() {
			mx.Lock() //lock
			total2++
			mx.Unlock() //unlock
		}()
	}

	// 3. Dengan Atomic - aman untuk operasi sederhana
	var total3 int64
	for range 10000 {
		// dengan atomic, kita bisa buat counter yg aman race condition dengan sederhana
		// tapi atomic ini terbatas hanya untuk variabel
		atomic.AddInt64(&total3, 1) // artinya kita bisa meningkatkan angka total secara atomik
		// total = total + 1
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("tidak melakukan locking, hasil `total`: %d \n", total1)
	fmt.Printf("melakukan locking dengan mutex. hasil  `total`: %d \n", total2)
	fmt.Printf("melakuakn locking dengan atomic, hasil total: %d \n", total3)
}
