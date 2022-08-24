package main

import (
	"fmt"
)

/*
	Problems:
	1. JustString and HugeString points to the same memory
		-> gc can't release unused HugeString tail
		-> if we change HugeString, JustString will be changed too

	2. String contains runes, slice expression work with bytes -> after cutting a rune can be divided

	*3. In general the shown example doesn't have problems, because it does nothing :)
*/

func main() {
	hugeString := CreateHugeString(10000)
	justString := CopyFirstNRunes(5, &hugeString)
	hugeString = ""
	fmt.Println(justString)
}

func CreateHugeString(n int) string {
	matrix := []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдежзийклмнопрстуфхцчшщъыьэюяѐ")

	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = matrix[i%len(matrix)]
	}

	return string(res)
}

func CopyFirstNRunes(n int, source *string) string {
	var res []rune
	i := 0
	for _, r := range *source {
		if i == n {
			break
		}

		res = append(res, r)
		i++
	}

	return string(res)
}
