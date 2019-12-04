package main

import "fmt"

// Vertex is two points
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {

	fmt.Println("Maps")
	{
		m = make(map[string]Vertex)
		m["Bell Labs"] = Vertex{
			40.68433, -74.39967,
		}
		fmt.Println(m["Bell Labs"])
	}

	fmt.Println("\nMap literals")
	{
		// Maps can be created via literals, keys are required
		var m = map[string]Vertex {
			"Bell Labs": Vertex{
				40.68433, -74.39967,
			},
			"Google": Vertex{
				37.42202, -122.08408,
			},
		}

		fmt.Println(m)  // 'map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]'

		// Value type can be omitted (if named type)
		m = map[string]Vertex{
			"Bell Labs": {40.68433, -74.39967},  // Type Vertex and ctor is implicit
			"Google":    {37.42202, -122.08408},
		}

		fmt.Println(m)  // 'map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]'
	}

	fmt.Println("\nMutating maps")
	{
		m := make(map[string]int)

		m["Answer"] = 42  // Set (and insert)
		fmt.Println("The value:", m["Answer"])
	
		m["Answer"] = 48  // Set (and replace)
		fmt.Println("The value:", m["Answer"])
	
		delete(m, "Answer")  // Delete
		fmt.Println("The value:", m["Answer"])
	
		v, ok := m["Answer"] // map provides two results (value, and existance) - go's version of 'keyExists'
		fmt.Println("The value:", v, "Present?", ok)	// 'The value: 0 Present? false'	

		// That's nice - Go won't scream if the key is invalid, but also wont require a 2nd call to get it's value
	}
}
