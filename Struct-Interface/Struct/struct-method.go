//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Circle struct { // struct dechlare
//	x, y, r float64
//}
//type Rectangele struct {
//	height        int
//	weight        int
//	inside_circle Circle
//}
//
//func Area(c Circle) float64 { // fucntion
//	return math.Pi * c.r * c.r
//}
//func (c *Circle) area2() float64 { // method
//	return math.Pi * c.r * c.r
//}
//func main() {
//	//--------------------------way of creating instance of struct--------------------
//	var c Circle // it can be accessable using Dot (.)
//	_ = c
//	c1 := Circle{0, 0, 0}
//	c1.x, c1.y, c1.r = 12, 13, 14
//	_ = c1
//	c2 := new(Circle)
//	_ = c2
//	//--------------------------- way to creating slice of struct----------------------
//	sl := []Circle{}
//	temp := Circle{5, 6, 7}
//	sl = append(sl, temp)
//	_ = sl
//	//-----------------------------------------------------------------------
//	temp = Circle{1, 3, 5}
//	fmt.Println("Area Using Function: ", Area(temp))
//	//..........................Method......................................
//	fmt.Println("Area Using Method: ", temp.area2())
//	//--------------------------Embedded Types------------------------------
//	box := Rectangele{12, 34, temp} // here, temp circle is inside of box rectangle
//	fmt.Println(box)
//
//}
