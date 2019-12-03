package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	
	// Pointers
	{
		i, j := 42, 36

		p := &i         // point to i
		fmt.Println(p)	// Pointer value (the value of p is the address of i) - (e.g. '0xc00008e000')
		fmt.Println(*p) // read i through the pointer
		
		*p = 21         // set i through the pointer
		fmt.Println(i)  // see the new value of i
	
		p = &j         // point to j
		*p = *p / 37   // divide j through the pointer
		fmt.Println(j) // see the new value of j
	}

	// Pointers to structs
	{
		v:= Vertex{1, 2}
		p := &v
		p.X = 1e9  // Can also be (*p).X - but the language allows shortcut

		fmt.Println(v)

		// This is where it gets confusing... printing a struct pointer shows that it's a reference.
		// Printing what it points to is the struct.
		fmt.Println(p)  // '&{1e9, 2}'
		fmt.Println(*p) // '{1e9, 2}'

		(*p).X = 2000
		fmt.Println(*p) // '{2000, 2}'
	}

	// Assigning a variable to equal a struct creates copy
	{
		type S struct {
			X int
		}

		v:= S{1}
		p := v  // Copy
		p.X = 2  // Does not alter v

		fmt.Println(p);
		fmt.Println(v);
	}

	// Struct Literals - behavior of properties upon instance of a struct
	{
		var (
			v1 = Vertex{1, 2}  // has type Vertex
			v2 = Vertex{X: 1}  // Y:0 is implicit
			v3 = Vertex{}      // X:0 and Y:0
			p  = &Vertex{1, 2} // has type *Vertex
		)
		
		fmt.Println(v1, p, v2, v3, *p);
	}
}