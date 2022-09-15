package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	dieChan := make(chan bool, 1) // buffer not to lock goroutine when its finished
	defer close(dieChan)

	wg.Add(1)
	go goroutineWithContext(ctx, &wg)
	go goroutineWithChannel(dieChan)

	time.Sleep(time.Second)
	dieChan <- true

	time.Sleep(time.Second)
	cancel()

	<-dieChan // wait for goroutineWithChannel() to finish working
	wg.Wait() // wait for goroutineWithContext() to finish working
}

func goroutineWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Started goroutine with context")

	<-ctx.Done() // wait for terminate signal
	// can use "select" if goroutine working with other channels

	fmt.Println("Finished goroutine with context")
}

func goroutineWithChannel(die chan bool) {
	fmt.Println("Started goroutine with channel")

	<-die // wait for terminate signal
	// also can use "select" if goroutine working with other channels

	die <- true // send signal back that goroutine finished working
	// can also use WaitGroup, like in goroutineWithContext()

	fmt.Println("Finished goroutine with channel")
}
