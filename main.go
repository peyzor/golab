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
}
