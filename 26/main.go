package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(AllSymbolsUnique("abcd"))
	fmt.Println(AllSymbolsUnique("abCdefAf"))
	fmt.Println(AllSymbolsUnique("aadcd"))
	fmt.Println(AllSymbolsUnique("АПРфЫы"))
	fmt.Println(AllSymbolsUnique("12341"))
}

func AllSymbolsUnique(str string) bool {
	m := make(map[rune]struct{})

	for _, ch := range str {
		if unicode.IsLetter(ch) {
			ch = unicode.ToLower(ch)
		}

		if _, ok := m[ch]; ok {
			return false
		}

		m[ch] = struct{}{}
	}

	return true
}
