package main

import "fmt"

func main() {
	// for unrepeatable elements
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{2, 5, 8, 12, 14, 26}

	m1 := make(map[int]struct{})
	for _, a := range arr1 {
		m1[a] = struct{}{}
	}

	m2 := make(map[int]struct{})
	for _, a := range arr2 {
		m2[a] = struct{}{}
	}

	var res []int
	for key := range m1 {
		if _, ok := m2[key]; ok {
			res = append(res, key)
		}
	}

	fmt.Println(res)
}
