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

func generics() {
	si := []int{1, 2, 3, 4, 5}
	result := Index(si, 3)
	fmt.Printf("first index appeared is: %v\n", result)

	si2 := []string{"holy", "moly", "xdd"}
	result = Index(si2, "moly")
	fmt.Printf("first index appeared is: %v\n", result)
}
