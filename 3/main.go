package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var res int

	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	m := sync.Mutex{}

	for _, a := range arr {
		go func(n int) {
			m.Lock()
			res += n * n
			m.Unlock()

			wg.Done()
		}(a)
	}

	wg.Wait()

	fmt.Println(res)
}
