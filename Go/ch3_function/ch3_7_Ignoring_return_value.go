// ASSIGNMENT
// In this code snippet, we only need our customer's first name. Ignore the last name so that the code compiles.
package main

import "fmt"

func ch3_7() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe"
}
