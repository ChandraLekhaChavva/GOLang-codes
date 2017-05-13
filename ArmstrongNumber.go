package main

import "fmt"
import "strconv"

func main() {
	for i := 0; i < 1000; i++ {
		if isArmStrong(i) {
			fmt.Println(i)
		}

	}
}

func isArmStrong(number int) bool {
	tmp := number
	digits := len(strconv.Itoa(number))
	sum := 0
	div := 0

	for i := 0; i < digits; i++ {
		div = tmp % 10
		sum = sum + (div * div * div)
		tmp = tmp / 10
	}
	return number == sum
}
