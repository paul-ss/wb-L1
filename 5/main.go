package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nSec := 5
	stop := time.After(time.Second * time.Duration(nSec))
	exchange := make(chan interface{})

	go func() {
		for {
			select {
			case <-stop:
				close(exchange)
				return
			default:
				exchange <- rand.Int()
				time.Sleep(time.Second)
			}
		}
	}()

	for data := range exchange {
		fmt.Println(data)
	}
}
