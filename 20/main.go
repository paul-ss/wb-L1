package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	line := "snow dog sun 123"
	fmt.Println(ReverseWordsOrder(line))
}

func ReverseWordsOrder(str string) string {
	sc := bufio.NewScanner(strings.NewReader(str))
	sc.Split(bufio.ScanWords)

	var strs []string
	for sc.Scan() {
		strs = append(strs, sc.Text())
	}

	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-1-i] = strs[len(strs)-1-i], strs[i]
	}

	return strings.Join(strs, " ")
}
