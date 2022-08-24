package main

import "fmt"

func main() {
	{
		a, b := 1, 2

		a = a ^ b
		b = b ^ a
		a = a ^ b

		fmt.Println(a, b)
	}

	{
		a, b := 1, 2

		a = a + b
		b = a - b
		a = a - b

		fmt.Println(a, b)
	}

	{
		a, b := 1, 2

		a = a * b
		b = a / b
		a = a / b

		fmt.Println(a, b)
	}
}
