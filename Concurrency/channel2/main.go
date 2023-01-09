package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	done := make(chan string)
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)

	fmt.Println("elapse : ", time.Since(start))
	fmt.Println("Finish Main Routine as well")
}
func task1(done chan string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Working: task 1")
	done <- "finsih task 1"
}
func task2(done chan string) {
	fmt.Println("working: task 2")
	done <- "finsih task 2"
}
func task3(done chan string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Working: task 3")
	done <- "finsih task 3"
}
func task4(done chan string) {
	fmt.Println("Working: task 4")
	done <- "finsih task 4"
}
