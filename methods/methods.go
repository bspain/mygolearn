package main

import (
	"fmt"
	"math"
)

// Vertex is an X and Y
type Vertex struct {
	X, Y float64
}

// Abs calculates the Absolute value of a Vertex
func (v Vertex) Abs() float64 {  // Go doesn't have classes, so instead, it uses a 'receiver' concept.
								// This function has a reciever of type Vertex named v

	return math.Sqrt(v.X*v.X + v.Y*v.Y)  // The recieved object (named v) can be referenced here
}

// MyFloat is a float64 
type MyFloat float64 //types need not be structs, they can be simple aliases of other types

// Absf calculates the Absolute value of a MyFloat
func (f MyFloat) Absf() float64 { // That allows for creation of method types
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Scale will scale a Vertex by f
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	fmt.Println("\nMethods")
	{
		v := Vertex{3, 4}
		fmt.Println(v.Abs())	// Abs() appears as a function on Vertex.  Better to think of this as an extension method	
	}

	fmt.Println("\nMethods continued")
	{
		f := MyFloat(-math.Sqrt2)
		fmt.Println(f.Absf())  // Absf() appears as a function on MyFloat

		// Extension methods w recievers only work on local types (those in the package), e.g. can't do this to types in other packages
		// including built in types like int
	}

	fmt.Println("\nPointer recievers")
	{
		// Normally methods that accept struct arguments will get a copy of the struct data.  A pointer is required if the function is to operate
		// on the struct data directly (pass by reference, not by value)
		v := Vertex{3, 4}
		v.Scale(10)		// Will modify the actual values of the vertex
		fmt.Println(v.Abs())	// '50'  ({3, 4} scaled by 10 { 30, 40 } -> Abs())
	}
}
