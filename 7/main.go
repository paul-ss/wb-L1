package main

import (
	"fmt"
	"sync"
)

func main() {
	const goNum = 500

	m := make(map[int]int)
	mux := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(goNum)
	for i := 0; i < goNum; i++ {
		go func(i int) {
			mux.Lock()
			defer mux.Unlock()
			m[i%20] = i

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(m)
}
