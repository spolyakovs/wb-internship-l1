package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// if there is only 1 goroutine, can use channel to send terminate signal
// but it's better to use context
func main() {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	workersNum := 3

	input := make(chan any, workersNum)
	defer close(input)

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-appSignal
		fmt.Printf("\nShutting down\n")
		close(appSignal)
		close(input)
		stop()
		os.Exit(0)
	}()

	for i := 0; i < workersNum; i++ {
		go workerWithContext(ctx, input)
	}

	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		input <- n
		time.Sleep(time.Second)
	}
}

// +: easy to use with any number of goroutines
// can use child context, context.WithTimeout, variables etc
func workerWithContext(ctx context.Context, input <-chan any) {
	for {
		select { // if both cases are ready, will choose RANDOM
		case <-ctx.Done():
			return
		case elem := <-input:
			fmt.Println(elem)
		}
	}
}
