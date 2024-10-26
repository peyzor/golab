package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//// In their uninitialized forms they are (nil)
	//var dict map[string]string
	////dict["x"] = "y" // can not assign to a nil map
	//var list []int
	//
	//fmt.Println(dict == nil)
	//fmt.Println(list == nil)
	//
	//// map literal
	//newDict := map[string]string{
	//	"x": "y",
	//}
	//newDict["y"] = "z"
	//
	//// if a key is not found the default value is going to be returned
	//val := newDict["not-found"]
	//fmt.Println(val == "")
	//
	//// the comma ok syntax for maps
	//if _, ok := newDict["not-found"]; !ok {
	//	fmt.Println("does not exist")
	//}
	//
	//dict1 := map[int]int{
	//	1:  2,
	//	10: 5,
	//	2:  3,
	//}
	//dict2 := map[int]int{
	//	10: 5,
	//	2:  3,
	//	1:  2,
	//}
	//// can hack map comparison like below
	//first := fmt.Sprintf("%#v", dict1)
	//second := fmt.Sprintf("%#v", dict2)
	//fmt.Println(first == second)

	//// map declaration (uninitialized)
	//var myMap2 map[string]int
	//_ = myMap2
	//
	//// map literal (initialized)
	//myDict := map[string]int{
	//	"x": 1,
	//	"y": 2,
	//	"z": 3,
	//}
	//
	//fmt.Printf("%#v\n", myDict)
	//for k := range myDict {
	//	delete(myDict, k)
	//	//delete(myDict, k) // deleting a non-existing key is a noop
	//}
	//
	//fmt.Printf("%#v\n", myDict)
	//
	//// make (initialized)
	//myMap := make(map[string]int)
	//myMap["x"] = 1
	//myMap["y"] = 2
	//myMap["z"] = 3
	//
	//fmt.Printf("%#v\n", myMap)

	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		input := in.Text()
		fmt.Printf("line no %d: %s\n", lines, input)
		lines++
	}

	if err := in.Err(); err != nil {
		fmt.Println(err)
	}
}
