package sunny

import "fmt"

func even_increment() func() int {
	x := 0
	return func() int {
		p := x + 2
		x = p
		return p
	}
}
func main() {
	c3 := even_increment()
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", c3())
	}

}
