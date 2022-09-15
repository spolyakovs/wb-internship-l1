package main

import "fmt"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temp := range temps {
		key := int(temp/10) * 10

		if _, ok := groups[key]; !ok {
			groups[key] = make([]float64, 0)
		}

		groups[key] = append(groups[key], temp)
	}

	fmt.Println(groups)
}
