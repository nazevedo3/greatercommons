// sort the []user by

// age
// last
// Also sort each []string “Sayings” for each user

// print everything in a way that is pleasant

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

	fmt.Println(users)

	// your code goes here

	sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	// fmt.Println("By age:", users)

	sort.Slice(users, func(i, j int) bool { return users[i].Last < users[j].Last })
	// fmt.Println("By name:", users)

	for i, v := range users {
		fmt.Println("\nPerson #", i+1)
		fmt.Println("\nUser First:", v.First, "\nUser Last:", v.Last, "\nUser Age:", v.Age)
		for i, val := range v.Sayings {
			fmt.Println("Saying #", i+1, val)

		}
	}
}
