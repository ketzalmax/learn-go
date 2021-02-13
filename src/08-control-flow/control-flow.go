package main

import "fmt"

func printSlice(label string, states []string) {
	fmt.Printf("%s len=%d cap=%d %v\n", label, len(states), cap(states), states)
}

func main() {

	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice Cream`, `Sunsets`},
	}

	m["Jaws"] = []string{`steel teeth`, `biting things`, `choking people`}

	for k, v := range m {
		fmt.Println("This is a record for", k, v)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	fmt.Println("------------------")
	delete(m, `no_dr`)

	for k, v := range m {
		fmt.Println("This is a record for", k, v)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	// simplemap := map[string]string{"name": "Ketzal", "eye color": "blue", "height": "6 feet"}
	// simplestringintmap := map[string]int{"age": 19, "heightinfeet": 6}

	// states := make([]string, 0, 57)

	// statesList := []string{`Alabama`, "Alaska", "American Samoa", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "District of Columbia", "Florida", "Georgia", "Guam", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Minor Outlying Islands", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Northern Mariana Islands", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Puerto Rico", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "U.S. Virgin Islands", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

	// for i := 0; i < len(statesList); i++ {
	// 	states = append(states, statesList[i])
	// 	fmt.Printf("cap %v, len %v, index %v, %v\n", cap(states), len(states), i, states[i])
	// }

	// for i, v := range states {
	// 	fmt.Println(i, v)
	// }
	// []string4, 45, 46, 47, 48, 49, 50, 51}
	// x = append(x[:3], x[6:]...)
	// x = append(x, 52)
	// fmt.Println(x)
	// x = append(x, 53, 54, 55)// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", x)
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", x)
}

// func main() {
// 	var a [5]int
// 	b := [5]int{41, 42, 43, 44, 45}
// 	fmt.Println("empty:", a)
// 	fmt.Println("populated:", b)

// 	c := [5]int{1, 2, 3, 4, 5}
// 	fmt.Printf("%T\n", c)
// 	fmt.Println("declared:", c)
// 	fmt.Println("populated:", b)
// }

// func main() {
// 	m := map[string]int{
// 		"James":       32,
// 		"Monkeypenny": 28,
// 	}

// 	fmt.Println(m["Monkeypenny"])
// 	valueOfJames, jerkwadOK := m["James"]
// 	fmt.Println(valueOfJames)
// 	fmt.Println("Is it OK to do this? ", jerkwadOK)
// }
