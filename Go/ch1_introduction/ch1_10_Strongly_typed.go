//ASSIGNMENT
// We'll be using simple basic authentication for the Textio API. This is how our users will communicate to us who they are and how many features they are paying for with each request to our API.

// The code on the right has a type error. Change the type of password to a string (but use the same numeric value) so that it can be concatenated with the username variable.
package main

import "fmt"

func ch1_10() {
	var username string = "wagslane"
	var password string = "20947382822"

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+password)
}
