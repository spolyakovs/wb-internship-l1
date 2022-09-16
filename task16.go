package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("Before: \n", slice)
	quicksort(slice)
	fmt.Println("\nAfter: \n", slice)
}

func quicksort(slice []int) {
	if len(slice) < 2 {
		return
	}

	pivotIndex := len(slice) - 1 // select last element as pivot to divide array
	leftMostIndex := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] < slice[pivotIndex] {
			// move everything less, than pivot to left side
			slice[leftMostIndex], slice[i] = slice[i], slice[leftMostIndex]
			leftMostIndex++
		}
	}

	// after cycle leftMostIndex will point to first elem, greater than pivot
	slice[leftMostIndex], slice[pivotIndex] = slice[pivotIndex], slice[leftMostIndex]

	quicksort(slice[:leftMostIndex])   // sort everything before
	quicksort(slice[leftMostIndex+1:]) // sort everything after
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(200) - 100 // random between -100 and 100
	}
	return slice
}
