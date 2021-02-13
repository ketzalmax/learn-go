package main

import "fmt"

// exercise 2
type cat struct {
	firstName string
	lastName  string
	favoriteCatFoodFlavors
}

type favoriteCatFoodFlavors []string

func main() {
	cat1 := cat{
		firstName:              "Fiesty",
		lastName:               "Var",
		favoriteCatFoodFlavors: favoriteCatFoodFlavors{`tuna`, `chicken`, `liver`},
	}
	cat2 := cat{
		firstName:              "Wilson",
		lastName:               "The Jerk",
		favoriteCatFoodFlavors: favoriteCatFoodFlavors{`tuna`, `dry`},
	}

	for i, v := range cat1.favoriteCatFoodFlavors {
		fmt.Println(i, v)
	}
	for _, v := range cat2.favoriteCatFoodFlavors {
		fmt.Println(v)
	}

	catMap := map[string]cat{
		cat1.lastName: cat1,
		cat2.lastName: cat2,
	}

	fmt.Println(catMap)

	for k, v := range catMap {
		fmt.Println(k, v)
		for i, v := range v.favoriteCatFoodFlavors {
			fmt.Println(i, v)

		}
	}

}

// Exercise 3
// type vehicle struct {
// 	doors int
// 	color string
// }

// type truck struct {
// 	vehicle
// 	fourWheel bool
// }

// type sedan struct {
// 	vehicle
// 	luxury bool
// }

// var truck1 = truck{
// 	vehicle: vehicle{
// 		doors: 2,
// 		color: "Blue",
// 	},
// 	fourWheel: true,
// }

// func main() {
// 	truck2 := truck{
// 		vehicle: vehicle{
// 			doors: 2,
// 			color: "White",
// 		},
// 		fourWheel: true,
// 	}
// 	sedan1 := sedan{
// 		vehicle: vehicle{
// 			doors: 4,
// 			color: "Red",
// 		},
// 		luxury: true,
// 	}

// 	fmt.Printf("%T\t%v\n", truck1, truck1)
// 	fmt.Printf("%T\t%v\n", truck2, truck2)
// 	fmt.Printf("%T\t%v\n", sedan1, sedan1)
// 	fmt.Printf("%T\t%v\n", sedan1.vehicle, sedan1.vehicle)
// 	fmt.Println("is sedan1 a luxury vehicle: ", sedan1.luxury)
// }

// Exercise 4
func main() {

	anonStruct := struct {
		simpleField   string
		embeddedSlice []string
		embeddedMap   map[string]int
	}{
		simpleField:   "simple field",
		embeddedSlice: []string{"slice 1", "slice 2", "slice 3"},
		embeddedMap:   map[string]int{"map1": 1, "map2": 2},
	}

	fmt.Println(anonStruct.simpleField)
	fmt.Println(anonStruct.embeddedSlice[0])
	fmt.Println(anonStruct.embeddedMap["map1"])
	fmt.Println()
	fmt.Println()
	fmt.Println()

}
