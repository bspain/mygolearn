package main

import (
	"fmt"
	"math"
)

// Abser interface
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println("\nImplicit interfaces");
	{
		var i I = T{"hello"};
		i.M()	// i is the I interface, and T implements I (by means of the M() extension method) - no where however, is it explicit that T implements I
				// e.g. there is no 'implements' or 'extends' keyword.
	}

}

// MyFloat is a float64
type MyFloat float64

// Abs calculates the absolute value of a MyFloat
func (f MyFloat) Abs() float64 {  // Valid implementation of Abser being applied to MyFloat
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex is an X, Y
type Vertex struct {
	X, Y float64
}

// Abs calculates the absolute value of a *Vertex
func (v *Vertex) Abs() float64 {  // Valid implementation of Abser being applied to *Vertex
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// Implicit interfaces

// I interface
type I interface {
	M()
}

// T is a string
type T struct {
	S string
}

// M prints the string of a T
func (t T) M() {
	// This method means type T implements the interface I,
	// but we don't need to explicitly declare that it does so.

	fmt.Println(t.S)
}
