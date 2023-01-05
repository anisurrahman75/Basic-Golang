package main

import "fmt"

func main() {
	//------------------------------------------------------ ArrAY
	//var arr [5]int // array dechlaration:
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} // sort way for Array Dechlaration
	for i, value := range arr {
		fmt.Println(i, value)
	}
	fmt.Printf("\n")
	//-------------------------------------------------------Slices
	//var sl1 []int
	sl1 := make([]int, 0) // using make function where require length of slice
	sl1 = append(sl1, 1, 2, 3, 4)
	fmt.Println(sl1)
	sl1 = append(sl1, 5)
	fmt.Println(sl1)
	sl1 = arr[:]
	fmt.Println(sl1)
	sl1 = sl1[0:5]
	fmt.Println(sl1)
	//------------------------------------------------------Map
	x := make(map[string]int) // map creation
	x["key"] = 10
	fmt.Println(x["key"], "map length; ", len(x))
	//delete(x, "key")
	fmt.Println(x["key"]) // return 0 if key isn't in map
	x["key"] = 0
	// check in map value exist or not
	if name, ok := x["key"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("NOT FIND")
	}
	element := make(map[string]map[string]string)
	_ = element
	element["sunny"] = map[string]string{
		"HomeTown": "dhaka",
		"Univ":     "Daffodil",
	}
	mp := make(map[string][]int) // create map of string slice
	mp["sunny"] = append(mp["sunny"], 5)
	fmt.Println(mp["sunny"])

	//Problems:
	y := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	mn := y[0]
	for _, value := range y {
		if value < mn {
			mn = value
		}
	}
	fmt.Println("minimum is: ", mn)

}
