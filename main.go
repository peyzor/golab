package main

import "fmt"

type Vertex struct {
	x, y float64
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// there are 2 special behaviours here for convenience
	// 1. struct fields can be accessed through a struct pointer
	v := Vertex{1, 2}
	p := &v
	p.x = 10
	//(*p).x = 10 // the language automatically does this for us (automatic dereferencing)
	fmt.Printf("%+v\n", p)

	// 2. pointer receivers take either value or a pointer
	v = Vertex{1.5, 2.25}
	v.Scale(5)
	//(&v).Scale(5) // the language automatically does this for us (automatic referencing)
	fmt.Printf("%+v\n", v)
	//ScaleFunc(v, 5)
	ScaleFunc(&v, 5)
	fmt.Printf("%+v\n", v)

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

	// a nil slice has length and capacity of 0 and has no underlying array
	var myNilSlice []int
	fmt.Println(myNilSlice, len(myNilSlice), cap(myNilSlice), myNilSlice == nil)

	// the make function allocates a zeroed array and return a slice that refers to that array
	mySlice := make([]int, 0, 5)
	fmt.Println(mySlice, len(mySlice), cap(mySlice), mySlice == nil)

	// the default is zero for lower bound and the length of the slice for the high bound
	mySlice = mySlice[:]
	fmt.Println(mySlice, len(mySlice), cap(mySlice), mySlice == nil)

	// but we can go beyond the length of the slice and into the underlying array
	// notice that we can't go beyond this unlike python it is a runtime error
	mySlice = mySlice[:cap(mySlice)]
	fmt.Println(mySlice, len(mySlice), cap(mySlice), mySlice == nil)

	f := fibonacci()
	for i := 1; i < 10; i++ {
		fmt.Println(f())
	}
}
