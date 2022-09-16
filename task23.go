package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sliceLen := 5
	slice := generateSlice(sliceLen)

	rand.Seed(time.Now().UnixNano())
	removeIndex := rand.Intn(sliceLen)

	fmt.Println("Slice: \n", slice)
	slice, _ = removeFromSlice(slice, removeIndex)
	fmt.Printf("\nSlice after removing slice[%v]: %v\n", removeIndex, slice)
	_, err := removeFromSlice(slice, -1)
	fmt.Println("\nError after removing slice[-1] (not exist):", err)
	_, err = removeFromSlice(slice, 30)
	fmt.Println("\nError after removing slice[30] (not exist):", err)
}

func removeFromSlice(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return slice, errors.New("index out of range")
	}

	return append(slice[:index], slice[index+1:]...), nil
}

// from task 16
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(200) - 100 // random between -100 and 100
	}
	return slice
}
