package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	p := append(s, 20)
	display(p)
	p = append(p[0:3], 10)
	display(p)
	p = p[0:cap(p)]
	display(p)
	s = append(s[0 : len(s)-1])
	display(s)

	s = append(s, 5)
	//s = append(s[0 : len(s)-1])
	display(s)
}
func display(s []int) {
	fmt.Printf("Slice length: %v\nSlice cap: %v\nSlice: %v", len(s), cap(s), s)
	fmt.Println()
	fmt.Println()

}
