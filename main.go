package main

import "fmt"

func main() {
	// In their uninitialized forms they are (nil)
	var dict map[string]string
	//dict["x"] = "y" // can not assign to a nil map
	var list []int

	fmt.Println(dict == nil)
	fmt.Println(list == nil)

	// map literal
	newDict := map[string]string{
		"x": "y",
	}
	newDict["y"] = "z"

	// if a key is not found the default value is going to be returned
	val := newDict["not-found"]
	fmt.Println(val == "")

	// the comma ok syntax for maps
	if _, ok := newDict["not-found"]; !ok {
		fmt.Println("does not exist")
	}

	dict1 := map[int]int{
		1:  2,
		10: 5,
		2:  3,
	}
	dict2 := map[int]int{
		10: 5,
		2:  3,
		1:  2,
	}
	// can hack map comparison like below
	first := fmt.Sprintf("%#v", dict1)
	second := fmt.Sprintf("%#v", dict2)
	fmt.Println(first == second)
}
