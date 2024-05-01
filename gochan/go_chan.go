package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// Fix 2: Use a loop body variable
		i := i // "i" shadows "i" from the for loop
		go func() {
			fmt.Println(i) // i from line 14
		}()
		/* Fix 1: Use a parameter
		go func(i int) {
			fmt.Println(i)
		}(i)
		*/
		/* BUG: ALL goroutine use the "i" for the for loop
		go func() {
			fmt.Println(i) // "i" from line 12
		}()
		*/
	}

	time.Sleep(10 * time.Millisecond)

	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch // receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got:", msg)
	}

	/* for/range does this
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("got: ", msg)
	}
	*/

	msg = <-ch // ch is closed
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <-ch // ch is closed
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)

	// ch <- "hi" // ch is closed -> panic

	fmt.Println(sleepSort([]int{100, 200, 300, 400}))
}

func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		n := n
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}

	var out []int

	for range values {
		n := <-ch
		out = append(out, n)
	}

	return out
}

/*
For every value "n" in values, spin a goroutine that will
- sleep "n" milliseconds
- Send "n" over a channel

In the function body, collect values from the channel to a slice and return it
*/

/*
- end & receive will block until opposite operation (*)
- receive from a closed channel will return the zero value without blocking
- send to a closed channel will panic
- closing a closed channel will panic
- send/receive to a nil channel will block forever
*/
