package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go task1()
	go task2()
	go task3()
	go task4()
	time.Sleep(110 * time.Millisecond)
	// since main function didn't wait for finish other's parallel function works, its stop execute.
	// if we can wait main function then the all other function will be execute.
	// this is why time slepp for delay
	fmt.Println("elapse : ", time.Since(start))

}
func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1")

}
func task2() {
	fmt.Println("task 2")
}
func task3() {

	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 3")
}
func task4() {
	fmt.Println("task 4")
}
