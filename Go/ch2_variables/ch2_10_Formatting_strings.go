// ASSIGNMENT
// Create a new variable called msg on line 9. It's a string that contains the following:

// Hi NAME, your open rate is OPENRATE percent
// Copy icon
// Where NAME is the given name, and OPENRATE is the openRate rounded to the nearest "tenths" place.
package main

import "fmt"

func ch2_10() {
	const name = "Saul Goodman"
	const openRate = 30.5

	// ?
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	// don't edit below this line

	fmt.Println(msg)
}
