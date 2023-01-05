package main

import "fmt"

func average(arr []int, n int) float64 { //pass by value
	var avg float64
	for _, value := range arr {
		avg += float64(value)
	}
	return float64(avg / float64(len(arr)))
}
func add(arr ...int) int { // Multiple  parameter
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func main() {
	var x []int
	for i := 0; i <= 10; i++ {
		x = append(x, i)
	}
	fmt.Println(average(x, len(x)))
	fmt.Println(add(1, 3), add(1, 2, 3)) // multiple paremeter don't support method overriding
	// ----------------Closure function-------------
	nextEven := makeEvenGenerator()

	fmt.Println(nextEven()) // 0
	//fmt.Println(nextEven()) // 2
	//fmt.Println(nextEven()) // 4
}
