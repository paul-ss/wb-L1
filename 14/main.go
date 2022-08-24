package main

import "fmt"

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan int)
	arr := []interface{}{1, "123", false, ch1, ch2, int64(123)}

	for _, a := range arr {
		DetectType(a)
	}

}

func DetectType(i interface{}) {
	switch val := i.(type) {
	case int:
		fmt.Printf("int type: %d\n", val)
	case string:
		fmt.Printf("string type: %s\n", val)
	case bool:
		fmt.Printf("bool type: %t\n", val)
	case chan interface{}:
		fmt.Printf("chan interface{} type: %p\n", val)
	default:
		fmt.Printf("unknown type: %v\n", val)
	}
}
