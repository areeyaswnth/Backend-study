// ASSIGNMENT
// We can speed up our contact-info lookups by using a map! Looking up a value in a map by its key is much faster than searching through a slice.

// Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and potentially an error. A user struct just contains a user's name and phone number.

// If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".

// The first name in the names slice matches the first phone number, and so on.
package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// ?
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	m := make(map[string]user)
	for i, n := range names {
		m[n] = user{name: n, phoneNumber: phoneNumbers[i]}
	}
	return m, nil
}

// don't touch below this line

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}
