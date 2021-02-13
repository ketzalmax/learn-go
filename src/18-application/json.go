package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// must uppercase/export to use JSON Marshal and Unmarshal
// could uppercase just field names
type ColorGroup struct {
	ID int
	// json tags: use `json:"key"` to set name key on JSON output
	// add omitempty to omit from struct if value is empty, e.g, json:"myName, omitempty"`
	// to omit a field in output use `json:"-"``
	// to generate a field with name "-" use `json:"-,"`
	Name   string `json:"ColorName"`
	Colors []string
}

// turn a string of JSON objects into a byte array in one step
var colorGroupsBlob = []byte(`[
	{"ID": 1, "Name": "Reds", "Colors": ["Crimson","Red","Ruby","Maroon"]},
	{"ID": 2, "Name": "Blues", "Colors": ["Navy","Cerulean","Sky","Indigo"]}
]`)

func main() {

	cgroup1 := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	cgroup2 := ColorGroup{
		ID:     1,
		Name:   "Blues",
		Colors: []string{"Navy", "Cerulean", "Sky", "Indigo"},
	}

	// Marshal: take a struct and convert to a JSON object, as an array of bytes
	b, err := json.Marshal(cgroup1)
	if err != nil {
		fmt.Println("ERROR! ", err)
	}
	// use os.Stdout to write out b
	os.Stdout.Write(b)

	b, err = json.MarshalIndent(cgroup2, "", "   ")
	if err != nil {
		fmt.Println("ERROR!! ", err)
	}
	// use fmt.printLn to write out b
	fmt.Println("\nINDENTED\n", string(b))

	// Create a byte array from a string in two steps and Marshal it
	var s string = `{"ID": 3, "Name": "Yellows", "Colors": ["Lemon","Canary","Citrine","Chartreuse"]}`
	bs := []byte(s)

	// yellows := ColorGroup{}
	var yellows ColorGroup

	err = json.Unmarshal(bs, &yellows)
	if err != nil {
		fmt.Println("ERROR!! ", err)
	}
	fmt.Println("the solitary yellows struct struct: ", yellows)

	// Unmarshal more than one JSON object into an array of structs
	var colorgroups []ColorGroup

	err = json.Unmarshal(colorGroupsBlob, &colorgroups)
	if err != nil {
		fmt.Println("ERROR!! ", err)
	}
	fmt.Println("array of colorgroup structs: ", colorgroups)
	fmt.Println("first struct in the array: ", colorgroups[0])
}
