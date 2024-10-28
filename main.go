package main

import "fmt"

func main() {
	//var i interface{}
	// the empty interface can hold any value of any type
	var i any

	i = 42
	fmt.Printf("%v\n", i)
	i = "xd"
	fmt.Printf("%v\n", i)

	// type casting
	i = float64(10)
	fmt.Printf("(%v, %[1]T)\n", i)

	// a type assertion provides access to an interface value's underlying concrete type
	myFloat := i.(float64)
	fmt.Printf("(%v, %[1]T)\n", myFloat)

	s, ok := i.(string)
	fmt.Printf("(%v, %[1]T), %t\n", s, ok)

	v := i.(string) // panic
	fmt.Printf("(%v, %[1]T)\n", v)
}
