package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr = RemoveIthEl(arr, 0)
	fmt.Println(arr)
	arr = RemoveIthEl(arr, len(arr)-1)
	fmt.Println(arr)
	arr = RemoveIthEl(arr, 3)
	fmt.Println(arr)
}

func RemoveIthEl(arr []int, i int) []int {
	if i >= len(arr) || i < 0 {
		return arr
	}

	return append(arr[:i], arr[i+1:]...)
}
