package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cntOfChannel := 4 // here is total number of routine
	arrsz := 20       // make sure  arr size is multiply by channel size, here my channel size 4
	pre, post, k := 0, (arrsz / cntOfChannel), (arrsz / cntOfChannel)
	slc := []int{}
	for i := 0; i < arrsz; i++ {
		slc = append(slc, rand.Int()%50)
	}
	n_sum := 0
	for _, i := range slc {
		//fmt.Printf("%d ", i)
		n_sum += i
	}
	fmt.Println("Noraml Total sum of slice: ", n_sum)
	done := make(chan int, cntOfChannel) // crating channel variabale
	temp := 0
	start := time.Now()
	for i := post; i <= arrsz; i += k {
		go sum(slc[pre:i], done)
		pre += k
		//temp += <-done
	}
	for i := 0; i < cntOfChannel; i++ {
		temp += <-done
	}

	close(done)
	fmt.Println("Toal sum of slice: ", temp)
	fmt.Println("elapse : ", time.Since(start))
}

func sum(slc []int, done chan int) {
	//time.Sleep(500 * time.Millisecond)
	temp := 0
	for _, i := range slc {
		temp += i
	}
	done <- temp
}
