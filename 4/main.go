/*
	This method uses wg to wait for termination of all gorutines.
	But maybe I can only close the input chanel and don't use wg?
	No, in that case the long task won't be completed
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Pool struct {
	wNum  int
	input chan interface{}
	wg    sync.WaitGroup
}

func NewPool(wNum int) *Pool {
	return &Pool{
		wNum:  wNum,
		input: make(chan interface{}, wNum),
	}
}

func (p *Pool) Start() {
	p.wg.Add(p.wNum)

	for i := 0; i < p.wNum; i++ {
		go func(i int) {
			for data := range p.input {
				fmt.Printf("worker %d: %v\n", i, data)
			}

			fmt.Printf("worker %d stops\n", i)
			p.wg.Done()
		}(i)
	}
}

func (p *Pool) Stop() {
	close(p.input)
	p.wg.Wait()
}

func (p *Pool) Send(data interface{}) {
	p.input <- data
}

func main() {
	pool := NewPool(3)
	pool.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)

loop:
	for {
		select {
		case <-stop:
			break loop

		default:
			pool.Send(rand.Int())
			time.Sleep(time.Millisecond * 1000)
		}
	}

	pool.Stop()
}
