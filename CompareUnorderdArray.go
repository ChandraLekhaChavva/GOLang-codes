package main

import (
	"fmt"
	"sort"
)

func main() {
	c1 := []string{"chandu", "lekha", "chavva"}
	c2 := []string{"lekha", "chandu", "chavva"}
	c3 := []string{"chandra", "lekha", "chavva"}

	fmt.Println(IsEqual(c1, c2))
	fmt.Println(IsEqual(c1, c3))
}

func IsEqual(a1 []string, a2 []string) bool {
	sort.Strings(a1)
	sort.Strings(a2)
	if len(a1) == len(a2) {
		for i, v := range a1 {
			if v != a2[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
