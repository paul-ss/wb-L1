package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%0.64b\n", SetIthBit(0, 63, false))
	fmt.Printf("%0.64b\n", SetIthBit(31, 2, true))
}

func SetIthBit(val int64, i int, isZero bool) int64 {
	if i < 0 || i > 63 {
		return val
	}

	var mask int64 = 1 << i

	if !isZero {
		return val | mask
	}

	mask = math.MaxInt64 &^ mask

	return val & mask
}
