package main

import "fmt"

func main() {
	a := 10
	b := 35

	fmt.Printf("Before: a: %v, b: %v\n", a, b)

	a, b = b, a

	fmt.Printf("After: a: %v, b: %v\n", a, b)
}
