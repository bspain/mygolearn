package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Decoding json to object is called 'unmarshalling'

	fmt.Println("Simple JSON into object")
	{
		type Message struct {
			Name string `json:"name"`
			Time int64  `json:"time"`
		}

		b := []byte(`
		{
			"name" : "Alice",
			"time" : 1294706395881547000
		}`)

		var m Message

		err := json.Unmarshal(b, &m)

		fmt.Printf("m: %v\nerr: %v", m, err)
	}

	fmt.Println("\n\nJSON Array into object")
	{
		type List struct {
			Entries []string `json:"entries"`
		}

		b := []byte(`
		{
			"entries" : [ "one", "two", "three" ]
		}`)

		var l List
		err := json.Unmarshal(b, &l)

		fmt.Printf("l: %v\nerr: %v", l, err)
	}

	fmt.Println("\n\nJSON Array of objects into object - converts to map")
	{
		type List struct {
			Entries []interface{} `json:"entries"`
		}

		b := []byte(`
		{
			"entries" : [ { "foo": "bar" }, { "baz" : "blue" }, { "pat" : "pot" } ]
		}`)

		var l List
		err := json.Unmarshal(b, &l)

		fmt.Printf("l: %v\nerr: %v", l, err) // 'l: {[map[foo:bar] map[baz:blue] map[pat:pot]]}'
	}

	fmt.Println("\n\nJSON map into object")
	{
		type MyMap struct {
			Enteries map[string][]string `json:"entries"`
		}

		b := []byte(`
		{
			"entries" : { 
				"foo" : [ "foo1" ], 
				"bar" : [
					"bar1",
					"bar2"
				], 
				"baz" : []
			}
		}`)

		var m MyMap
		err := json.Unmarshal(b, &m)

		fmt.Printf("m: %v\nerr: %v", m, err) // 'm: {map[bar:[bar1 bar2] baz:[] foo:[foo1]]}'
	}

	fmt.Println("\n\nJSON mapping to nested object")
	{
		// We have to do some JSON-splunking and reflection to get this object to align with the Unmarsheled object
		type MyEntries struct {
			entries []float64
		}

		b := []byte(`
		{
			"entries" : [ 2, 3, 4 ]
		}`)

		var unk map[string]interface{}
		err := json.Unmarshal(b, &unk)

		fmt.Printf("unk: %v\nerr: %v\n", unk, err)            // 'unk: map[entries:[2 3 4]]'
		fmt.Printf("unk[\"entries\"]: %v\nerr: %v\n", unk["entries"], err) // 'unk["entries"]: [2 3 4]'

		l := unk["entries"].([]interface{})  // l is a 'generic' array
		fmt.Printf("l: %v\nerr: %v\n", l, err) // 'l: [2 3 4]'

		// There is no easy way to convert the generic []interface{} into []float64, so manually walking the values it is :(
		var myEntries MyEntries
		for _, e := range l {
			myEntries.entries = append(myEntries.entries, e.(float64))
		}
		
		fmt.Printf("myEntries: %v\nerr: %v\n", myEntries, err) // 'unTypedEntries: map[entries:[2 3 4]]'
	}
}
