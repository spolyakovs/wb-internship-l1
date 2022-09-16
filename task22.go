package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a, b int64 = rand.Int63(), rand.Int63()

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println("a + b:", add(a, b))
	fmt.Println("a - b:", subtract(a, b))
	fmt.Println("a * b:", multiply(a, b))
	fmt.Println("a / b:", divide(a, b))
}

func add(a, b int64) *big.Int {
	res := big.NewInt(0)
	return res.Add(big.NewInt(a), big.NewInt(b))
}

func subtract(a, b int64) *big.Int {
	res := big.NewInt(0)
	return res.Sub(big.NewInt(a), big.NewInt(b))
}

func multiply(a, b int64) *big.Int {
	res := big.NewInt(0)
	return res.Mul(big.NewInt(a), big.NewInt(b))
}

func divide(a, b int64) *big.Int {
	res := big.NewInt(0)
	// panics if b == 0 !!!
	return res.Div(big.NewInt(a), big.NewInt(b))
}
