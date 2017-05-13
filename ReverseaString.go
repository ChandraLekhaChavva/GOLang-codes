package main

import (
	"fmt"
)

func main() {
	strrev := "ChandraLekha"
	fmt.Println(strrev)
	strrev = reverseString(strrev)
	fmt.Println(strrev)
}

func reverseString(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}
