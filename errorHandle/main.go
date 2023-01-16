package main

import (
	"errors"
	"fmt"
	"os"
)

func Division(a int, b int) (float64, error) {

	if b == 0 {
		return float64(0), errors.New("Can't Divide by Zero")
	} else {
		return (float64(a) / float64(b)), nil
	}
}
func main() {
	path := "./file.txt"
	text, fileError := os.ReadFile(path)
	if fileError != nil {
		fmt.Println("Not Find any file")
	} else {
		fmt.Println(string(text))
	}
	value, divError := Division(7, 5)
	if divError != nil {
		fmt.Println(divError)
	} else {
		fmt.Println("Division ans: ", value)
	}
}
