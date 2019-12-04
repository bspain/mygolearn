package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {  // c chanel for an int must be predefined (via make)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	fmt.Println("Routines")
	{
		// Don't expect predictable output here, each thread is being scheduled independently
		go say("world")
		say("hello")
	}

	fmt.Println("\nChannels")
	{
		// <- is used to send and recieve values through a channel (sync/locked memory space)
		s := []int{7, 2, 8, -9, 4, 0}

		c := make(chan int)

		go sum(s[len(s)/2:], c)  // Create one routine to process the first half of s
		go sum(s[:len(s)/2], c)  // Create a 2nd routine to process the second half

		x, y := <-c, <-c // recieve from channel c  (This is blocking until work is complete)
							

		fmt.Println(x, y, x+y)

		// Can also think of the channel as a queue, this additional call to the channel will error because
		// there are no more goroutines processing work
		// fmt.Println(<-c)
	}

	fmt.Println("\nRange and Close")
	{
		fibonacci := func (n int, c chan int) {
			x, y := 0, 1
			for i := 0; i < n; i++ {
				c <- x
				x, y = y, x+y
			}

			close(c)  // Only the sender should close a channel, never the receiver.  (Sending on a closed channel with panic)
		}

		c := make(chan int, 10)
		go fibonacci(cap(c), c)

		// for-range on a channel, will continue to iterate on the channel, doing <- until it is closed
		for i := range c {
			fmt.Println(i)
		}
	}

	fmt.Println("\nSelect")
	{
		fibonacci := func (data, stop chan int) {  // c channel will get data until stop has been 'signaled' -- Allows the recieve to incidate it's done reading from c
			x, y := 0, 1

			for {
				select {	// This will block until one of the cases is valid - e.g. data can recieve, or stop can be written to
				case data <- x:
					x, y = y, x+y
				case <- stop:
					fmt.Println("Stopping")
					return
				}
			}
		}

		var c = make(chan int)
		var q = make(chan int, 1)

		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(<-c)
			}
			q <- 0
		}()

		fibonacci(c, q)
	}
}