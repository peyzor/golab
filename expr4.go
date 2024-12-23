package main

import (
	"fmt"
	"io"
	"strings"
)

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("and the names is: %s with age: %d", p.name, p.age)
}

func expr4() {
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

	//v := i.(string) // panic
	//fmt.Printf("(%v, %[1]T)\n", v)

	// type switches allow several type assertions in series
	var w any
	w = int64(1)
	switch v := w.(type) {
	case int:
		fmt.Printf("it is int: %d\n", v)
	case string:
		fmt.Printf("it is string: %s\n", v)
	default:
		fmt.Printf("(%v, %[1]T)\n", v)
	}

	a := Person{"Arthur Morgan", 44}
	fmt.Println(a)

	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)
	n := copy(dst, src)
	fmt.Println("Copied elements: ", n)
	fmt.Println("Destination slice ", dst)
}
