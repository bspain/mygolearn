package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	rand.Seed( int64(time.Now().Nanosecond()) );
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))


	// Capital letters are exported, lower case letters are not
	fmt.Println(math.Pi);

	// Types always follow variables, this makes it easy to write closures
	// Just read l to r , sum is the result of a function that takes two int's, returns and int, with implementation or a+b auto-invoked with (3,4)
	sum := func(a, b int) int { return a+b } (3, 4);
	fmt.Println("Sum is", sum); // 7


	// Functions can return multiple results
	a, b := func(x, y string) (string, string) { return y, x } ("world", "hello");
	fmt.Println(a, b) // Hello world

	// Functions can have 'named' arguments.  A 'naked' return then, would return the named args
	split := func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return;
	}

	fmt.Println(split(13)) // 5 8
	fmt.Println(split(17)) // 7 10

	// Var can be used to declare variables (both local and global scope)
	var i int = 1;
	var c, python, java = true, false, "no!";

	// := is just implicit var (but it only works in local scope - can't be used outside a function)
	// := also applies implicit typing.
	j := 4;
	
	fmt.Println(i, j, c, python, java);

	// The expression T(v) converts the value v to the type T
	// Go requires explicit conversions
	{
		i := 42;
		var f float64 = float64(i);
		u := uint(f);

		fmt.Println(i, f, u);
	}

	// Type inference occurs when the right side contains an untyped numeric constant
	// could be int, float64, complex128...
	{
		i := 42           // int
		f := 3.142        // float64
		g := 0.867 + 0.5i // complex128

		fmt.Printf("i is of type %T\n", i);
		fmt.Printf("f is of type %T\n", f);
		fmt.Printf("g is of type %T\n", g);
	}
}
