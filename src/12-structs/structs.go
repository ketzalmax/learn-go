package main

import "fmt"

type weapons struct {
	primaryWeapon   string
	secondaryWeapon string
}

type person struct {
	weapons
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			weapons: weapons{
				primaryWeapon:   "Walther PPK",
				secondaryWeapon: "Exploding Pen",
			},
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	var sa2 secretAgent
	sa2.person.first = "Pinky"
	sa2.last = "The Good"
	sa2.age = 5
	sa2.ltk = true

	p1 := person{
		weapons{
			primaryWeapon:   ".380",
			secondaryWeapon: "Poisoned Hairpin",
		},
		"Miss",
		"Moneypenny",
		27,
	}

	// anonymous struct
	anonP := struct {
		evilName  string
		evilPet   string
		evilLaugh bool
	}{
		evilName:  "Dr. No",
		evilPet:   "cat",
		evilLaugh: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)
	fmt.Println(anonP)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.primaryWeapon)
	// full path to nested struct fields
	fmt.Println(sa1.person.weapons.primaryWeapon, sa1.person.weapons.secondaryWeapon)
	// promoted nested struct field
	fmt.Println("promoted nested struct field:", sa1.primaryWeapon)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(anonP.evilName, anonP.evilPet, anonP.evilLaugh)

	fmt.Printf("\n%T\n", p1)

}
