package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	weight float64
}
type Circle struct {
	x       float64
	y       float64
	radious float64
}

func (r Rectangle) area() float64 {
	return r.weight * r.height
}
func (c Circle) area() float64 {
	return math.Pi * c.radious * c.radious
}

func main() {
	// when all of struct's have same method then we can use interface,
	// let's say here a slice  of different types of struct but one common name of method, then we can make a interface
	// above Rectangle and Circle both have common method as area , but implementation different.
	r := Rectangle{10, 12}
	c := Circle{12, 24, 12.4}

	printArea(r)
	printArea(c)
}

func printArea(s Shape) {
	fmt.Println(s.area())
}
