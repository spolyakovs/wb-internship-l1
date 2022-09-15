package main

import "fmt"

type Set map[string]bool // lets assume, that sets are maps like this

func (s Set) Put(key string) {
	s[key] = true
}

func (s Set) Remove(key string) {
	delete(s, key)
}

func main() {
	set1 := make(Set)
	set2 := make(Set)

	set1.Put("a")
	set1.Put("a")
	set1.Put("b")
	set1.Put("c")

	set2.Put("b")
	set2.Put("c")
	set2.Put("d")
	set2.Put("e")

	fmt.Println(getIntersection(set1, set2))
}

func getIntersection(set1, set2 map[string]bool) map[string]bool {
	res := make(Set)
	for key := range set1 {
		if set1[key] == true && set2[key] == true {
			res.Put(key)
		}
	}
	return res
}
