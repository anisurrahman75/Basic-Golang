package main

import "fmt"

type Numbers []interface{}

func (n *Numbers) appendLast(v Numbers) {
	*n = append(*n, v...)
}
func (n *Numbers) appendFirst(v Numbers) {
	*n = append(v, *n...)
}
func (n *Numbers) popLast() {
	*n = append((*n)[0:len(*n)-1], Numbers{}...)
}
func (n *Numbers) popFirst() {
	*n = append((*n)[1:len(*n)], Numbers{}...)
}
func ready(v interface{}) Numbers {
	x := Numbers{v}
	return x
}
func deque() {
	s := Numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s)
	x := ready("Anisur")
	s.appendFirst(x)
	fmt.Println("Appending ", x, "Into The First")
	fmt.Println(s)
	x = ready("Sunny")
	s.appendLast(x)
	fmt.Println("Appending ", x, "Into The Last")
	fmt.Println(s)
	s.popLast()
	fmt.Println("Pop Last Element")
	fmt.Println(s)
	s.popFirst()
	fmt.Println("Pop First Element")
	fmt.Println(s)
}
func main() {
	//deque()

}
