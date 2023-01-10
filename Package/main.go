package main

import (
	"fmt"
	v "github.com/anisurrahman75/Basic-Golang/Package/math"
	"math/rand"
)

func main() {
	fmt.Println("Simple interest calculation")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := v.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
	// Max..Min function from Math Function
	slc := []int{}
	for i := 0; i < 10; i++ {
		slc = append(slc, rand.Int()%50)
	}
	fmt.Println("All sclice Elements: ", slc)
	fmt.Println("Minimum of this Slice is: ", v.Min(slc))
	fmt.Println("Maximum of this Slice is: ", v.Max(slc))

}
