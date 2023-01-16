package main

import (
	"fmt"
	"strconv"
)

type customeError struct {
	name    string
	errCode int
}

func (c customeError) Error() string {
	return c.name + "\nError Code: " + strconv.Itoa(c.errCode)
}
func Division(a int, b int) (float64, error) {
	if b == 0 {
		return 0.0, customeError{"Divide By Zero Not Possible", -1}
	} else {
		return float64(a) / float64(b), nil
	}
}
func main() {
	value, x := Division(10, 0)
	if x != nil {
		fmt.Println(x.Error())
	} else {
		fmt.Println(value)
	}
}

//Youtube: https://www.youtube.com/watch?v=VMveb4GqRck
