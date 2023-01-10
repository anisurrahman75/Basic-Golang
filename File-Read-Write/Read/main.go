package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./create-file1.txt"
	databyte, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
