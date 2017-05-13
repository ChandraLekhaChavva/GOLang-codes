package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "chandra works smart"
	fmt.Println(reverseWordsInString(str))
}

func reverseWordsInString(str string) (finalStr string) {
	strSlice := strings.Split(str, " ")
	strLen := len(strSlice)

	// Is length required
	for i := 0; i < strLen/2; i++ {
		tmp := strSlice[strLen-i-1]
		strSlice[strLen-i-1] = strSlice[i]
		strSlice[i] = tmp
	}
	str = strings.Join(strSlice, " ")
	return str
}
