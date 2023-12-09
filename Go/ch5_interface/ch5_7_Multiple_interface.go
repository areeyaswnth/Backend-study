// ASSIGNMENT
// Complete the required methods so that the email type implements both the expense and printer interfaces.

// COST()
// If the email is not "subscribed", then the cost is 0.05 for each character in the body. If it is, then the cost is 0.01 per character.

// Return the total cost of the entire email.

// PRINT()
// The print method should print to standard out the email's body text
package main

import (
	"fmt"
)

func (e email) cost() float64 {
	// ?
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	} else {
		return float64(len(e.body)) * 0.01
	}
}

func (e email) print() {
	// ?
	println(e.body)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test2(e expense2, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================\n")
}

func ch5_7() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test2(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test2(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test2(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test2(e, e)
}
