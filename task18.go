package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	sync.RWMutex
	value uint64
}

func (c *counter) Increment() {
	c.Lock()
	c.value++
	c.Unlock()
}

func (c counter) GetValue() uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.value
}

func main() {
	// to better test this, run with "--race" flag
	c := &counter{}
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			c.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter with lock:", c.GetValue())

	// also can use atomic
	var value uint64 = 0

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&value, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter with atomic:", value)
}
