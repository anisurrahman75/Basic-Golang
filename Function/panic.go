package sunny

import "fmt"

func sillySusan() {
	fmt.Println("silly susan called")
	panickingPeter()
	fmt.Println("silly susan finished")
}

func panickingPeter() {
	fmt.Println("panicking peter called")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("now way  i recover function are here ")
		}
	}()
	panic("Hackers:  end of everything!")
	fmt.Println("panicking peter finished")
	//}

	//func main() {
	fmt.Println("cascading panics")

	//sillySusan()
	fmt.Println("Also not Print this! ")
}
