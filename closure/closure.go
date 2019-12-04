package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// functions will retain variables within the closure (these are part of the stack)
	// so subsequent calls to this function can act upon them
	prev := 0
	curr := 1
	calls := 0
	return func() int {
		// This is an internal function, but still has access to the closure surrounding it (so prev, curr, calls are in scope)
		if calls == 0 {
			calls++
			return 0
		} else if calls == 1 {
			calls++
			return 1
		} else {
			var temp = prev
			prev = curr
			curr = temp + prev
			return curr
		}
	}
}

func fibonacciWithDefer() func() int {
	prev := 0	
	curr := 1
	calls := 0

	// Good use of defer here, because we want to return the current number of calls, but then ensure calls is incremented
	incrementCalls := func() { calls++ }

	return func() int {
		if calls == 0 || calls == 1 {
			defer incrementCalls() // increment calls after return
			return calls
		}

		var temp = prev
		prev = curr
		curr = temp + prev
		return curr
	}
}


func main() {
	f := fibonacci()
	f2 := fibonacciWithDefer()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
		fmt.Println(f2())
	}
}