package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// special stop chan
	{
		stop := make(chan interface{})
		someCh := make(chan interface{})

		wg.Add(1)
		go func() {
			for {
				select {
				case <-stop:
					fmt.Println("stopped")
					wg.Done()
					return
				case <-someCh:
				}
			}
		}()

		stop <- 1
	}

	// close chanel
	{
		someCh := make(chan interface{})
		wg.Add(1)
		go func() {
			for _ = range someCh {

			}

			fmt.Println("stopped")
			wg.Done()
		}()

		close(someCh)
	}

	// context - cancel
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		wg.Add(1)
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("stopped")
					wg.Done()
					return
				}
			}
		}()

		cancel()
	}

	// context - timeout
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		wg.Add(1)
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("stopped")
					wg.Done()
					return
				}
			}
		}()
	}

	// lol
	{
		m := sync.Mutex{}
		m.Lock()

		wg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			fmt.Println("stopped")
			wg.Done()
		}()

		m.Unlock()
	}

	wg.Wait()
}
