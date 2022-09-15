package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// very similar to task2
	// difference only in handleOutput() (added lines 46 and 50, changed line 48)
	array := [...]int{2, 4, 6, 8, 10}

	squareChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(len(array))

	for _, elem := range array {
		go getSquare(&wg, elem, squareChan)
	}

	doneChan := make(chan bool) // needed to wait for checkOutput()
	// otherwise main() will finish before and last element won't be printed

	go handleOutput(squareChan, doneChan)

	wg.Wait()
	close(squareChan) // ends "range" in line 44
	<-doneChan
}

func getSquare(wg *sync.WaitGroup, v int, squareChan chan<- int) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000) + 1000
	time.Sleep(time.Duration(n) * time.Millisecond) // sleep 1-2 sec

	squareChan <- v * v
}

func handleOutput(squareChan <-chan int, doneChan chan<- bool) {
	sum := 0
	for square := range squareChan {
		sum += square
	}
	fmt.Println(sum)
	doneChan <- true
}
