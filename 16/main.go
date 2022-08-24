package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//{
	//	arr := []int{2, 5, 2, 5, 8, 7, 5, 3, 2, 1, 34, 5, 43, 2, 32, 4, 4, 6, 11}
	//	QuickSort(arr)
	//	fmt.Println(arr)
	//}
	//
	//{
	//	arr := []int{5, 4, 3, 2, 1}
	//	QuickSort(arr)
	//	fmt.Println(arr)
	//}
	//
	for i := 0; i < 100; i++ {
		fmt.Printf("%d, ", rand.Int()%100)
	}
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotId := partition(arr)
	QuickSort(arr[:pivotId])
	QuickSort(arr[pivotId+1:])
}

func partition(arr []int) int {
	if len(arr) < 2 {
		panic("len(arr) < 2")
	}

	//if len(arr) == 2 {
	//	if arr[1] < arr[0] {
	//		arr[1], arr[0] = arr[0], arr[1]
	//	}
	//
	//	return 0
	//}

	l, r := 0, len(arr)-1

	pivotId := getPivotId(arr, l, r)
	arr[r], arr[pivotId] = arr[pivotId], arr[r]
	pivotId = r

	for i, j := l, r-1; ; /*i < j*/ {
		for arr[i] < arr[pivotId] && i < j {
			i++
		}

		for arr[j] >= arr[pivotId] && i < j {
			j--
		}

		if i == j {
			if arr[i] >= arr[pivotId] {
				arr[i], arr[pivotId] = arr[pivotId], arr[i]
				return i
			}

			return pivotId
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println("ops")
	return 0
}

func getPivotId(arr []int, r, l int) int {
	return l
}
