package main

import "fmt"

type PipeFunc func(in <-chan int, out chan<- int)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	funcs := []PipeFunc{WriteDigits(arr), Mul2, Out}

	in := make(chan int, 5)
	for _, f := range funcs {
		out := make(chan int, 5)

		go func(f PipeFunc, in, out chan int) {
			f(in, out)
			close(in)
		}(f, in, out)

		in = out
	}
}

func WriteDigits(arr []int) func(in <-chan int, out chan<- int) {
	return func(in <-chan int, out chan<- int) {
		for _, a := range arr {
			out <- a
		}
	}
}

func Mul2(in <-chan int, out chan<- int) {
	for a := range in {
		out <- a * 2
	}
}

func Out(in <-chan int, out chan<- int) {
	for a := range in {
		fmt.Println(a)
	}
}
