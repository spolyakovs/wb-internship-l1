package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Now:", time.Now().String())
	sleep(3)
	fmt.Println("After sleep:", time.Now().String())
}

func sleep(seconds int) {
	start := time.Now()
	for {
		if time.Since(start) >= time.Duration(seconds)*time.Second {
			break
		}
	}
}
