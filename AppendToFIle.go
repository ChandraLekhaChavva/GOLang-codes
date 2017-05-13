package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("could not open")
	}
	defer file.Close()
	_, err = file.WriteString("Chandra is coding in GOlang")
	if err != nil {
		fmt.Println("Could not write to file")
	} else {
		fmt.Println("Append to file")
	}
}
