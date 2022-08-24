package main

import (
	"fmt"
)

func main() {
	s := Set{}

	for i := 0; i < 300000; i++ {
		s.Add(fmt.Sprintf("%d", i))
	}

	fmt.Println(len(s.DFS()))
	s.Delete("1234")

	for i := 0; i < 300000; i++ {
		if !s.Exists(fmt.Sprintf("%d", i)) {
			fmt.Printf("key '%d' not exists\n", i)
		}
	}

	for i := 0; i < 300000; i++ {
		s.Delete(fmt.Sprintf("%d", i))
	}

	fmt.Println(s.DFS())
}

type Node struct {
	l, r *Node
	val  string
}

type Set struct {
	root *Node
}

func (s *Set) Add(key string) {
	var add func(n **Node)
	add = func(n **Node) {
		if *n == nil {
			*n = &Node{val: key}
			return
		}

		if (*n).val > key {
			add(&(*n).l)
		} else if (*n).val < key {
			add(&(*n).r)
		}
	}

	add(&s.root)
}

func (s *Set) Delete(key string) {
	found := search(key, &s.root)
	if found == nil {
		return
	}

	if (*found).r == nil {
		*found = (*found).l
		return
	}

	curr := &(*found).r
	// todo fix
	for (*curr).l != nil {
		curr = &(*curr).l
	}
	min := *curr
	*curr = (*curr).r

	min.r = (*found).r
	min.l = (*found).l
	*found = min
}

func (s *Set) Exists(key string) bool {
	return search(key, &s.root) != nil
}

func search(key string, n **Node) **Node {
	if *n == nil {
		return nil
	}

	if key < (*n).val {
		return search(key, &(*n).l)
	} else if key > (*n).val {
		return search(key, &(*n).r)
	} else {
		return n
	}
}

func (s *Set) DFS() (res []string) {
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}

		dfs(n.l)
		res = append(res, n.val)
		dfs(n.r)
	}

	dfs(s.root)
	return
}
