package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "главрыба"
	fmt.Println(ReverseString1(str))
	fmt.Println(ReverseString2(str))
}

func ReverseString1(str string) string {
	in := []rune(str)
	out := make([]rune, len(in))

	for i := len(in) - 1; i >= 0; i-- {
		out[len(in)-1-i] = in[i]
	}

	return string(out)
}

func ReverseString2(str string) string {
	sb := strings.Builder{}

	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		sb.WriteRune(r)
		str = str[:len(str)-size]
	}

	return sb.String()
}

