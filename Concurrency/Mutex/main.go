package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func mutexIncrement(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x) // wrong ans for running concurrently.
	x = 0
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go mutexIncrement(&w, &m)
	}
	w.Wait()
	fmt.Println("Using Mutex final value of x", x) // wrong ans for running concurrently.
}
