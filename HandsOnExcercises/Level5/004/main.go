//Create and use an anonymous struct.

package main

import "fmt"

func main() {

	as := struct {
		first      string
		last       string
		occupation string
		hobbies    []string
		friends    map[string]int
	}{
		first:      "Roger",
		last:       "Wilco",
		occupation: "Janitor",
		hobbies:    []string{"Space travel", "living", "Avoiding the Vohaul"},
		friends: map[string]int{
			"Pirates of Pestulon": 444,
			"Sludge Vohaul":       555,
		},
	}

	fmt.Println(as)
	fmt.Println(as.last)

	for k, v := range as.friends {
		fmt.Println(k, v)
	}

	for i, val := range as.hobbies {
		fmt.Println(i, val)
	}

}
