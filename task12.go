package main

import "fmt"

type Set map[string]bool // from task11

func (s Set) Put(key string) {
	s[key] = true
}

func (s Set) Remove(key string) {
	delete(s, key)
}

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	res := make(Set)
	for _, elem := range s {
		res.Put(elem)
	}

	fmt.Println(res)
}
