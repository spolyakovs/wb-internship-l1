package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	timeout := 5

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	var wg sync.WaitGroup

	dataChan := make(chan int)
	defer close(dataChan)

	wg.Add(2)

	go writeData(ctx, &wg, dataChan)
	go readData(ctx, &wg, dataChan)

	wg.Wait()
}

func writeData(ctx context.Context, wg *sync.WaitGroup, dataChan chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		data := rand.Intn(100)

		select {
		case <-ctx.Done():
			wg.Done()
			return
		case dataChan <- data:

		}

		time.Sleep(time.Second)
	}
}

func readData(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case data := <-dataChan:
			fmt.Println(data)
		}
	}
}
