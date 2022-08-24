package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	fmt.Println("Random order")
	RandomOrder(arr)

	fmt.Println("Correct order")
	CorrectOrder(arr)
}

func RandomOrder(arr []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	m := sync.Mutex{}

	for _, a := range arr {
		go func(n int) {
			m.Lock()
			fmt.Println(n * n)
			m.Unlock()

			wg.Done()
		}(a)
	}

	wg.Wait()
}

func CorrectOrder(arr []int) {
	res := make([]int, len(arr))

	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	for id, a := range arr {
		go func(id, n int) {
			res[id] = n * n
			wg.Done()
		}(id, a)
	}
	wg.Wait()

	for _, a := range res {
		fmt.Println(a)
	}
}
