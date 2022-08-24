package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for _, a := range arr {
		key := int(a) / 10 * 10
		m[key] = append(m[key], a)
	}

	for key, a := range m {
		fmt.Printf("%d: %v\n", key, a)
	}
}
