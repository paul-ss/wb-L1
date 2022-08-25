package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 62)
	b := big.NewInt(1 << 61)

	fmt.Println(big.NewInt(1).Add(a, b).String())
	fmt.Println(big.NewInt(1).Add(big.NewInt(1).Neg(b), a).String())
	fmt.Println(big.NewInt(1).Div(a, b).String())
	fmt.Println(big.NewInt(1).Mul(a, b).String())
}
