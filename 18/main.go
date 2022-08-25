package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Incr(1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(c.Show())
}

type Counter struct {
	val int64
}

func (c *Counter) Incr(a int64) int64 {
	return atomic.AddInt64(&c.val, a)
}

func (c *Counter) Show() int64 {
	return atomic.LoadInt64(&c.val)
}
