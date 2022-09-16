package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sliceLen := 20
	slice := generateSlice(sliceLen)
	quicksort(slice)

	rand.Seed(time.Now().UnixNano())
	findIndex := rand.Intn(sliceLen)

	fmt.Println("Slice: \n", slice)
	fmt.Printf("\nFind (slice[%v]): %v\n", findIndex, binarySearch(slice, slice[findIndex]))
	fmt.Println("\nFind 300 (not exist):", binarySearch(slice, 300))
	fmt.Println("\nFind -300 (not exist):", binarySearch(slice, -300))
}

func binarySearch(slice []int, value int) int {
	// slice should be SORTED

	// if value > last element
	if value > slice[len(slice)-1] {
		return -1
	}

	// if went here, than slice[0] <= value <= slice[len(slice)-1]
	leftMostIndex, rightLeastIndex := 0, len(slice)-1
	for leftMostIndex < rightLeastIndex {
		i := (leftMostIndex + rightLeastIndex) / 2
		switch {
		case slice[i] >= value:
			rightLeastIndex = i
		case slice[i] < value:
			leftMostIndex = i + 1
		}
	}

	// if value < slice[0] or simply not there
	if value < slice[leftMostIndex] {
		return -1
	}

	return leftMostIndex
}

// from task 16
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

// from task 16
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(200) - 100 // random between -100 and 100
	}
	return slice
}
