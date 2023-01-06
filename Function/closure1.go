package sunny

import "fmt"

func print_age(age int) {
	fmt.Println("My Current age is: ", age)
}
func main() {
	// anonymous function are used in closure:
	// anonymous function be  like:
	func() {
		fmt.Println("Hello my  Name is Anisur Rahman ")
	}() // function called from here.
	c1 := func(add string) {
		fmt.Println("My Hometown: ", add)
	}
	c1("Gazipur")
	// closure outside main function: age printing
	c2 := print_age
	c2(20)

}
