package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	Id     int
	status string
}

func generateOrder(n int) []*Order {
	orders := []*Order{}

	for i := range n {
		order := Order{
			Id:     i + 1,
			status: "pending",
		}
		orders = append(orders, &order)
	}

	return orders
}

func processOrder(orderChan <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()
	// loop berikut menerima value dari channel terus menerus hingga ditutup.
	for order := range orderChan {
		time.Sleep(time.Duration(rand.Intn(500)+10) * time.Millisecond)
		status := []string{"delivered", "shipped"}[rand.Intn(2)]
		order.status = status
		fmt.Printf("procssing order: %d with status: %s\n", order.Id, order.status)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	n := 20
	orderChan := make(chan *Order, n)

	go func() {
		wg.Done()
		for _, order := range generateOrder(n) {
			orderChan <- order
		}
		// karena kita mengirim data ke channel pada function ini, maka kita dapat panggil fungsi close di sini
		close(orderChan)
		fmt.Printf("done generating order\n")
	}()

	go processOrder(orderChan, &wg)

	wg.Wait()
	fmt.Printf("All operation is done!\n")
}
