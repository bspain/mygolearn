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

	fmt.Println("\n\nJSON Array into array of numbers - float64 is base number type")
	{
		type List struct {
			Entries []float64 `json:"entries"`
		}

		b := []byte(`
		{
			"entries" : [ 2, 3, 4 ]
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

	fmt.Println("\n\nJSON mapping of an unknown object")
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

		fmt.Printf("unk: %v\nerr: %v\n", unk, err)                         // 'unk: map[entries:[2 3 4]]'
		fmt.Printf("unk[\"entries\"]: %v\nerr: %v\n", unk["entries"], err) // 'unk["entries"]: [2 3 4]'

		l := unk["entries"].([]interface{})    // l is a 'generic' array
		fmt.Printf("l: %v\nerr: %v\n", l, err) // 'l: [2 3 4]'

		// There is no easy way to convert the generic []interface{} into []float64, so manually walking the values it is :(
		var myEntries MyEntries
		for _, e := range l {
			myEntries.entries = append(myEntries.entries, e.(float64))
		}

		fmt.Printf("myEntries: %v\nerr: %v\n", myEntries, err) // 'myEntries: {[2 3 4]}'
	}

	fmt.Println("\n\nJSON mapping of nested object")
	{
		// This object is within the JSON structure, but we don't want to type out the whole thing
		type MyEntries struct {
			Entries []float64	`json:"entries"`
		}

		b := []byte(`
		{
			"resp": {
				"list": {
					"entries" : [ 2, 3, 4 ]
				}
			}
		}`)

		// Treat the Json as a generic map[string] nested set to get to what we want
		var unk map[string]map[string]interface{}
		err := json.Unmarshal(b, &unk)

		fmt.Printf("err: %v\n", err)
		fmt.Printf("unk: %v\n", unk)										// 'unk: map[resp:map[list:map[entries:[2 3 4]]]]'
		fmt.Printf("unk[\"resp\"]: %v\n", unk["resp"])						// 'unk["resp"]: map[list:map[entries:[2 3 4]]]'
		fmt.Printf("unk[\"resp\"][\"list\"]: %v\n", unk["resp"]["list"])	// 'unk["resp"]["list"]: map[entries:[2 3 4]]'

		// Re-marshal list back out to JSON
		l, err2 := json.Marshal(unk["resp"]["list"])						

		fmt.Printf("l: %v\nerr2: %v\n", string(l), err2)					// 'l: {"entries":[2,3,4]}'

		// Now, unmarshal list as the JSON we care about
		var myEntries MyEntries
		err3 := json.Unmarshal(l, &myEntries)

		fmt.Printf("myEntries: %v\nerr3: %v\n", myEntries, err3) 			// 'myEntries: {[2 3 4]}'
	}

	fmt.Println("\n\nJSON decoding of string or string array property")
	{
		b := []byte(`
		{
			"groups" : [
				{
					"suites" : [ "one", "two", "three" ],
					"alias" : "foo"
				},
				{
					"suites" : [ "four", "five" ],
					"alias" : [ "bar", "baz" ]		
				}
			]
		}`)

		var l GroupList
		err := json.Unmarshal(b, &l)

		fmt.Printf("err: %v\n", err)
		fmt.Printf("l: %v\n", l)		// 'l: {[{[one two three] foo} {[four five] [bar baz]}]}'

		for i, v := range l.Groups {

			// Struct method handles getting alias always as a []string
			alias := v.GetAlias()

			fmt.Printf("Group# %v\nSuites: %v, Alias: %v\n", i, v.Suites, alias)

			/*
			Group# 0
			Suites: [one two three], Alias: [foo]
			Group# 1
			Suites: [four five], Alias: [bar baz]
			*/
		}
	}
}

// Types for JSON decoding example

// Group is a []string of suites and an alias (either string or []string)
type Group struct {
	Suites 		[]string				`json:"suites"`
	Alias		interface{}				`json:"alias"`		// Cound be string or []string
}

// GroupList is a list of groups
type GroupList struct {
	Groups 	[]Group		`json:"groups"`
}

// GetAlias converts the Alias of a Group into a []string
func (g *Group) GetAlias() (alias []string) {
	
	switch v := g.Alias.(type) {
	case string: 
		alias = append(alias, v)
	case []interface{}:
		for _, e := range v {
			alias = append(alias, e.(string))
		}
	}

	return
}