package main

import "fmt"

type Vertex struct {
	x int
	y int
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
}
