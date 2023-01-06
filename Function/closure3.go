package sunny

import "fmt"

func fib() func() int {
	a, b := 0, 1
	fmt.Printf("Fibonacchi series using closure: ")
	return func() int {
		x, y := a, b
		z := x + y
		a, b = b, z
		return z
	}
}
func main() {
	series := fib()
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", series())
	}
	println()
}
