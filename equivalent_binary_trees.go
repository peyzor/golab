package main

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walking(t *Tree, ch chan int) {
	defer close(ch)
	Walk(t, ch)
}

func Same(t1, t2 *Tree) bool {
	x := make(chan int)
	y := make(chan int)

	go Walking(t1, x)
	go Walking(t2, y)

	for {
		v1, ok1 := <-x
		v2, ok2 := <-y

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 {
			break
		}
	}
	return true
}
