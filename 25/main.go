package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	Sleep(time.Second)
	d := time.Since(t)
	fmt.Println("Slept for", d.Microseconds(), "microseconds")
}

func Sleep(d time.Duration) {
	ch := time.After(d)
	<-ch
}
