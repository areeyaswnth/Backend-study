// ASSIGNMENT
// Our Textio customers want to know how long they have had accounts with us.

// Follow the instructions in the comment provided. You will create a new variable called accountAgeInt that will be the truncated, integer version of accountAge.
package main

import "fmt"

func ch2_5() {
	accountAge := 2.6

	// create a new "accountAgeInt" here
	// it should be the result of casting "accountAge" to an integer
	accountAgeInt := int64(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
