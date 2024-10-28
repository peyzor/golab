package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

type Node[T any] struct {
	next *Node[T]
	val  T
}

func (n *Node[T]) print() {
	head := n
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}

func printFunc[T any](n *Node[T]) {
	head := n
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}

func generics() {
	si := []int{1, 2, 3, 4, 5}
	result := Index(si, 3)
	fmt.Printf("first index appeared is: %v\n", result)

	si2 := []string{"holy", "moly", "xdd"}
	result = Index(si2, "moly")
	fmt.Printf("first index appeared is: %v\n", result)

	x3 := Node[int]{nil, 3}
	x2 := Node[int]{&x3, 2}
	x1 := Node[int]{&x2, 1}
	x1.print()
	printFunc(&x1)
}
