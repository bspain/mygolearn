package main

import "fmt"

// Person is a name and age
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)  // This is akin to adding a 'toString()' method to Person.  It adheres to the 'Stringer' interface as defined in "fmt"
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	// Validate that a implements Stringer
	// TODO: Still seems like I'm missing something though... I need to 
	// declase i as fmt.Stringer to get that interface off of a... the next
	// check seems redundent then.  (Looks like 'reflect' package is designed for this specific case.)
	var i fmt.Stringer = a;
	_, ok := i.(fmt.Stringer)
	fmt.Printf("Does Person implement Stringer? %v", ok)  // Actually, it has to, or we would have errored out by now

	fmt.Println("\nExercise: Stringers")
	{
		hosts := map[string]IPAddr{
			"loopback":  {127, 0, 0, 1},
			"googleDNS": {8, 8, 8, 8},
		}
		for name, ip := range hosts {
			fmt.Printf("%v: %v\n", name, ip)
		}
	}
}


// Exercise, implementing a String() for IPAddr

// IPAddr is 4 bytes
type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}