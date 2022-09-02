package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (u ByAge) Len() int {
	return len(u)
}

func (u ByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u ByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

type ByLastName []user

func (u ByLastName) Len() int {
	return len(u)
}

func (u ByLastName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u ByLastName) Less(i, j int) bool {
	return u[i].Last < u[j].Last
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	// your code goes here

	fmt.Println("Idades antes do sort de idade")

	for i := 0; i < len(users); i++ {
		fmt.Println("\t", users[i].First, ":", users[i].Age)
	}

	sort.Sort(ByAge(users))
	fmt.Println("Idades depois do sort de idade")
	for i := 0; i < len(users); i++ {
		fmt.Println("\t", users[i].First, ":", users[i].Age)
	}

	fmt.Println("Last Name antes do sort de Last Name")
	for i := 0; i < len(users); i++ {
		fmt.Println("\t", users[i].Last)
	}

	sort.Sort(ByLastName(users))
	fmt.Println("Last Name depois do sort de Last Name")
	for i := 0; i < len(users); i++ {
		fmt.Println("\t", users[i].Last)
	}

	for i, user := range users {
		fmt.Printf("%v - %v %v:\n", i, user.First, user.Last)
		for i, currentSaying := range user.Sayings {
			fmt.Printf("\t%v - %v\n", i, currentSaying)
		}
	}

	for _, user := range users {
		sort.Strings(user.Sayings)
	}

	for i, user := range users {
		fmt.Printf("%v - %v %v:\n", i, user.First, user.Last)
		for i, currentSaying := range user.Sayings {
			fmt.Printf("\t%v - %v\n", i, currentSaying)
		}
	}

}

// @TODO:
// - Partindo do código abaixo, ordene os []user por idade e sobrenome.
//     - https://play.golang.org/p/BVRZTdlUZ_
// - Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
