package sunny

import "fmt"

func main() {

	fmt.Println("1")
	defer fmt.Println("2") // defer  are execute when main function end.
	defer func(string2 string) {
		fmt.Println("Hello my name is: ", string2)
	}("sunny")
	fmt.Println("3")
	fmt.Println("4")

}
