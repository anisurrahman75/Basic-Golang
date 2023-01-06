package sunny

import (
	"errors"
	"fmt"
)

func division(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("divided by zero")
		return 0, err
	}
	ret := a / b
	return ret, nil
}
func main() {
	result, err := division(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
