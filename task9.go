package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	array := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	inChan := make(chan int)
	outChan := make(chan int)
	doneSquareChan := make(chan bool)
	doneOutputChan := make(chan bool)

	go handleSquare(inChan, outChan, doneSquareChan)
	go handleOutput(outChan, doneOutputChan)

	for _, elem := range array {
		wg.Add(1)
		go func(elem int) {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(1000) + 1000
			time.Sleep(time.Duration(n) * time.Millisecond) // sleep 1-2 sec

			inChan <- elem
			wg.Done()
		}(elem)
	}

	wg.Wait()

	close(inChan)
	<-doneSquareChan
	close(outChan)
	<-doneOutputChan
}

func handleSquare(inChan <-chan int, outChan chan<- int, doneChan chan<- bool) {
	for elem := range inChan {
		outChan <- elem * elem
	}
	doneChan <- true
}

func handleOutput(outChan <-chan int, doneChan chan<- bool) {
	for elem := range outChan {
		fmt.Println(elem)
	}
	doneChan <- true
}
