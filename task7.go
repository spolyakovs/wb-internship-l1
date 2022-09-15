package main

import (
	"fmt"
	"strconv"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	data map[string]int
}

func MakeSafeMap() SafeMap {
	return SafeMap{
		data: make(map[string]int),
	}
}

func (m *SafeMap) Put(key string, value int) {
	m.Lock()
	m.data[key] = value
	m.Unlock()
}

func (m SafeMap) Get(key string) (value int, ok bool) {
	m.RLock()
	value, ok = m.data[key]
	m.RUnlock()
	return
}

func main() {
	data := MakeSafeMap()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			data.Put(key, i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()

			value, ok := data.Get(key)
			if !ok {
				fmt.Printf("data[%s] couldn't be accessed\n", key)
				return
			}
			fmt.Printf("data[%s]: %v\n", key, value)
		}(strconv.Itoa(i))
	}

	wg.Wait()
}
