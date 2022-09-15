package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s, i, b, c interface{} = "string", 10, true, make(chan int)

	printType(s)
	printType(i)
	printType(b)
	printType(c)
}

func printType(val interface{}) {
	// can use reflect
	fmt.Println("Reflect:", reflect.TypeOf(val))

	// or use %T in string format
	fmt.Printf("String format: %T\n", val)

	// or use type switch, but it can't detect generic channels, slices, maps etc
	// in theory can try all types with recursion, but can be problematic
	// for example map[map[string]struct]chan int
	fmt.Print("Type switch: ")
	switch v := val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown type,", v)
	}
	fmt.Println()
}
