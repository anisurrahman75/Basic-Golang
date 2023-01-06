package main

import "fmt"

func arr_pass_value(arr [4]int) {
	arr[0] = 5
}
func arr_pass_ref(arr *[4]int) {
	arr[0] = 5
}

func main() {
	// arr Pointer:
	arr := [4]int{1, 2, 3, 4}
	// value pass:
	fmt.Println("Before passing: ", arr)
	arr_pass_value(arr)
	fmt.Println("After passing passing: ", arr)

	// reference pass:
	fmt.Println("Before passing: ", arr)
	arr_pass_ref(&arr)
	fmt.Println("After passing passing: ", arr)
	// slice pointer: ( slice eventually a pointer so here is no pass by value only pass by ref,

}
