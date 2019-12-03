package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	// Basic loop syntax
	{
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}

		fmt.Println(sum) // 45

		// Loop variables are only visible in the context of the loop
		//	fmt.Println(i); -- 'undefined: i'
	}

	// for init; condition; post
	// init and post are optional (Go's version of a while loop)
	{
		sum := 1
		for sum < 1000 {
			sum += sum
		}

		fmt.Println(sum) // 1024
	}

	// if statement
	// normal if looks pretty standard
	if 1 < 2 { fmt.Println("true")}

	// if can declare scoped variable however (which is only visible in context of the if, and any companion else)
	pow := func (x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}

		// fmt.Println(v);   // Not allowed, 'undefined: v'

		return lim
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Switch in go , can declare local variable, automatically implies 'break'
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch without any condition (e.g. 'switch true') can be useful for long if-else chains
	{
		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
		}	
	}


	// Defer
	deferred := func () {
		// Will execute after the surrounding function completes (any call arguments are evaluated immediately)
		msg := "world"
		defer fmt.Println(msg)

		msg = "hello";
		fmt.Println(msg)

		// hello
		// world
	}
	deferred();

	// Defer is like pushing a function with values to the stack to be executed on completion.
	countdown := func() {
		defer fmt.Println("Blast OFF!")

		for i := 1; i <= 3; i++ {
			defer fmt.Println(i)
		}
	
		fmt.Println("Countdown!")
	}
	countdown();
}