package main

import "fmt"

type Myslice []int

func (s Myslice) remmoveIth(index int) []int {
	return append(s[:index], s[index+1:]...)
}
func main() {
	sl := Myslice{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(sl)
	sl = sl.remmoveIth(3)
	fmt.Println(sl)

}
