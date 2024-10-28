package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func fastFib(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	// only the sender must close a channel to indicate end of work
	close(c)
	//c <- 69 // send on a closed channel panics
}

func quitFib(c, quit chan int) {
	a, b := 0, 1

	for {
		select {
		case c <- a:
			a, b = b, a+b
		case <-quit:
			fmt.Println("quitting")
			return
		}
	}
}

func goroutines() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	var a chan int
	fmt.Printf("%v %[1]T\n", a)

	// if the buffer size is too small it causes a panic on send
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// channels are not like files; you don't usually need to close them.
	// closing is only necessary when the receiver must be told there are no more values coming
	// such as to terminate a range loop
	fibC := make(chan int, 10)
	go fastFib(20, fibC)

	// Mind Blown: we can range over channels
	for v := range fibC {
		fmt.Println(v)
	}
	fmt.Println(<-fibC) // receive on a closed channel
	x, ok := <-fibC     // it's also indicated with comma ok syntax
	fmt.Printf("%v, %v", x, ok)

	cFib := make(chan int)
	qFib := make(chan int)

	go quitFib(cFib, qFib)

	for i := 0; i < 10; i++ {
		fmt.Println(<-cFib)
	}
	qFib <- 0

	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)

	// the default case in select is run if no other case is ready
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick.")
	//	case <-boom:
	//		fmt.Println("BOOM!")
	//		return
	//	default:
	//		fmt.Println("    .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

	chx1 := make(chan string)
	chx2 := make(chan string)

	go func() {
		chx1 <- "Hello"
	}()

	go func() {
		chx2 <- "World"
	}()

	// we use select on its own without an enclosing "for" when we only need to process one thing
	// and then move on
	select {
	case msg1 := <-chx1:
		fmt.Println(msg1)
	case msg2 := <-chx2:
		fmt.Println(msg2)
	case <-time.After(1 * time.Nanosecond):
		fmt.Println("timed out!")
	}

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			chx1 <- "Ping"
		}
	}()

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			chx2 <- "Pong"
		}
	}()

	for {
		select {
		case msg1 := <-chx1:
			fmt.Println(msg1)
		case msg2 := <-chx2:
			fmt.Println(msg2)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Timeout! xd")
		}
	}
}
