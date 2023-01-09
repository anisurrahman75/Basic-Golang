package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { // another routine
		defer wg.Done() // execute after task1
		defer wg.Done() // execute after task2
		task1()
		task2()
	}()
	wg.Wait() // this function waiting untill wait grp finish
	fmt.Println("I am ending main function")
}
func task1() {
	time.Sleep(time.Second)
	fmt.Println("hello i am doing task1")
}
func task2() {
	time.Sleep(time.Second)
	fmt.Println("hello i am doing task2")
}
