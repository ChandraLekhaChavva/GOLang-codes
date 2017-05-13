package main

import (
	"fmt"
)

func main() {
	n := 0
	for {
		n = n + 1
		if n*n%1000000 == 269696 {
			break
		}

	}
	fmt.Println(n)
}
