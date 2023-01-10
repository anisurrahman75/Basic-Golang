package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Now try some File read write")
	text := "Hello I am sunny, Currently Learning Golang "
	file, err := os.Create("./create-file1.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	text_length, err := io.WriteString(file, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("length of this text is: ", text_length)
}
