package main

import "fmt"

// TODO change human with action

type Action struct {
	sentence string
}

func (a *Action) Speak() {
	fmt.Println(a.sentence)
}

type Human struct {
	Action
	sentence string
}

func (h *Human) Greet() {
	fmt.Println(h.sentence)
}

func main() {
	h := &Human{
		sentence: "Hi, I'm John",
		Action: Action{
			sentence: "Give me a cup of coffee, please",
		},
	}

	h.Speak()
}
