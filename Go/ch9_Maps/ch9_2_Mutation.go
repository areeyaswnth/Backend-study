// ASSIGNMENT
// It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

// Complete the deleteIfNecessary function.

// Check the scheduledForDeletion bool to determine if they are scheduled for deletion or not.

// If the user doesn't exist in the map, return the error not found.
// If they exist but aren't scheduled for deletion, return deleted as false with no errors.
// If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
// NOTE ON PASSING MAPS
// Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy.
package main

import (
	"errors"
	"fmt"
	"sort"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// ?
	elem, check := users[name]
	if check {
		if elem.scheduledForDeletion {
			delete(users, name)
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, errors.New("not found")
	}

}

// don't touch below this line

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted:", name)
		return
	}
	fmt.Println("Did not delete:", name)
}

func main() {
	users := map[string]user{
		"john": {
			name:                 "john",
			number:               18965554631,
			scheduledForDeletion: true,
		},
		"elon": {
			name:                 "elon",
			number:               19875556452,
			scheduledForDeletion: true,
		},
		"breanna": {
			name:                 "breanna",
			number:               98575554231,
			scheduledForDeletion: false,
		},
		"kade": {
			name:                 "kade",
			number:               10765557221,
			scheduledForDeletion: false,
		},
	}
	test(users, "john")
	test(users, "musk")
	test(users, "santa")
	test(users, "kade")

	keys := []string{}
	for name := range users {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	fmt.Println("Final map keys:")
	for _, name := range keys {
		fmt.Println(" - ", name)
	}
}
