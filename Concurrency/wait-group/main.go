package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { // another routine
		defer wg.Done() // execute after task1
		task1()
	}()
	go func() { // another routine
		defer wg.Done() // execute after task2
		task2()
	}()
	wg.Wait() // this function waiting untill wait grp finish
	fmt.Println("I am ending main function")
	fmt.Println("Total time: ", time.Since(start))
}
func task1() {
	time.Sleep(4 * time.Second)
	fmt.Println("hello i am doing task1")
}
func task2() {
	time.Sleep(5 * time.Second)
	fmt.Println("hello i am doing task2")
}
