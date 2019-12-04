package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		var a [2]string // An array of two strings

		a[0] = "Hello"
		a[1] = "World"

		fmt.Println(a[0], a[1]) // 'Hello World'  Println auto-formats comma separated values
		fmt.Println(a)          // '[Hello World]'  Array formatted output

		// Initialize array with literal
		primes := [6]int{2, 3, 5, 7, 11, 13}
		fmt.Println(primes)
	}

	// Slices - Arrays have a fixed size, Slices are dynamically sized, flexible views into an array (and are much more common)
	{
		var runes [6]rune // This is an array
		runes[0] = 'A'
		runes[1] = 'B'
		runes[2] = 'C'
		runes[3] = 'D'
		runes[4] = 'E'
		runes[5] = 'F'

		var s []rune = runes[1:4] // This is a 'slice' into (view of) the array.  Range is 'half-open' [ low : high ] - low is taken, high is omitted
		fmt.Println(s)

		fmt.Println(runes[0:1]) // 0th element - char('A') = 65
		fmt.Println(runes[5:6]) // 5th element - char('F') = 70

		//fmt.Println(primes[5:7]) // Error, slice out of bounds
	}

	// Slices can overlap, and modifying the slice modifyes the underlying data value

	{
		// Slice literal - creates an underlying array, and then provides a slice
		s := []struct {
			i int
			b bool
		}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
		}
		fmt.Println(s)

		// Even though s is a slice, r can still be a slice of it (r just becomes a slice of the underlying array for s)
		r := s[1:2]
		fmt.Println(r)

		s[1].i = 4
		s[1].b = true
		fmt.Println(r)
	}

	// Slice lenght (len) and capacity (cap)
	{
		printSlice := func(s []int) {
			fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
		}

		s := []int{2, 3, 5, 7, 11, 13}
		printSlice(s)

		// Slice the slice to give it zero length.
		s = s[:0]
		printSlice(s)

		// Extend its length.
		s = s[:4]
		printSlice(s)

		// Drop its first two values.
		s = s[2:]
		printSlice(s)

	}

	// Use 'make' to auto-generate a slice (with underlying array) zeroed out.
	{
		printSlice := func(s string, x []int) {
			fmt.Printf("%s len=%d cap=%d %v\n",
				s, len(x), cap(x), x)
		}

		a := make([]int, 5)
		printSlice("a", a)

		b := make([]int, 0, 5) //len = 0, capacity = 5
		printSlice("b", b)

		c := b[:2]
		printSlice("c", c)

		d := c[2:5]
		printSlice("d", d)
	}

	// Slices can contain any type, including other slices
	{
		// Create a tic-tac-toe board.
		board := [][]string{
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
		}

		// The players take turns.
		board[0][0] = "X"
		board[2][2] = "O"
		board[1][2] = "X"
		board[1][0] = "O"
		board[0][2] = "X"

		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
	}

	// append is used to dynamically grow a slice (will allocate more space for the underlying array)
	fmt.Println("\nAppending to a slice")
	{
		printSlice := func(s []int) {
			fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
		}

		var s []int // Allocate a 'nil' slice
		printSlice(s)

		s = append(s, 2048) // first arg is slice, 2nd is values to append
		printSlice(s)       // '[2048]'

		s = append(s, 3) // Slice, and underlying array grow as needed
		printSlice(s)    // ' [2048 3]'

		s = append(s, 9, 5, 2) // 2nd arg... is dynamic, and will add/allocate accordingly
		printSlice(s)
	}

	// range - a special way to get an iterator for a slice
	fmt.Println("\nRange")
	{
		var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

		for i, v := range pow {
			fmt.Printf("2**%d = %d\n", i, v)
		}

		// That's basically a nice alternative to
		for i := 0; i < len(pow); i++ {
			fmt.Printf("2**%d = %d\n", i, pow[i])
		}
	}

	fmt.Println("\nRange continued")
	{
		// Use _ to skip either the index, or the value
		var pow = make([]int, 10); // Slice of 10 0's

		// Use an index for everything in pow
		for i := range pow {
			pow[i] = 1 << uint(i) // e.g. bit shift, e.g. 2 * i
		}

		// Use the value for everything in pow (e.g. foreach)
		for _, v := range pow {
			fmt.Printf("%d\n", v)
		}
	}
}
