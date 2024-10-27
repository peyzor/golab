package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// struct fields can be accessed through a struct pointer
	v := Vertex{1, 2}
	p := &v
	p.x = 10
	//(*p).x = 10 // the language automatically does this for us
	fmt.Printf("%+v\n", p)

	names := [4]string{
		"John",
		"George",
		"Paul",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:]
	fmt.Println(a, b)

	// changing data in slice changes the data in the underlying array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// the length of a slice is the number of elements it contains
	// the capacity of a slice is the number of elements in the underlying array, counting from the first
	// element in the slice
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	// notice the capacity here
	s = s[2:]
	printSlice(s)
}
