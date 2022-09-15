package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	var data int64 = rand.Int63()
	var bitNum int64 = 4
	var bit uint8 = 0
	dataNew, ok := SetBitTo(data, bitNum, bit)
	if !ok {
		return
	}

	fmt.Println("Int before:", data)

	fmt.Printf("Int set %v bit to %v: %v\n", bitNum, bit, dataNew)

	dataNew, ok = SetBitTo(data, bitNum, 1-bit)
	if !ok {
		return
	}
	fmt.Printf("Int set %v bit to %v: %v\n", bitNum, 1-bit, dataNew)
}

func SetBitTo(data int64, bitNum int64, bit uint8) (int64, bool) {
	if bitNum > 63 {
		return data, false
	}

	switch bit {
	case 0:
		newData := data &^ (1 << bitNum)
		return newData, true
	case 1:
		newData := data | (1 << bitNum)
		return newData, true
	default:
		return data, false
	}
}
