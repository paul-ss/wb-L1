package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11, 12, 13, 14, 15, 16}

	{
		// searches left 10 id
		id := lSearch(0, len(arr)-1, func(params ...interface{}) bool {
			m := params[0].(int)
			return arr[m] >= 10
		})

		fmt.Println(id)
	}

	{
		// searches right 10 id
		id := rSearch(0, len(arr)-1, func(params ...interface{}) bool {
			m := params[0].(int)
			return arr[m] <= 10
		})

		fmt.Println(id)
	}
}

func lSearch(l, r int, check func(params ...interface{}) bool) int {
	for l < r {
		m := (l + r) / 2

		if check(m) {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func rSearch(l, r int, check func(params ...interface{}) bool) int {
	for l < r {
		m := (l + r + 1) / 2

		if check(m) {
			l = m
		} else {
			r = m - 1
		}
	}

	return l
}
