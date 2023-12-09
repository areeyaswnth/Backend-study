// ASSIGNMENT
// Fix the bug on line 12. The if statement should print "Message sent" if the messageLen is less than or equal to the maxMessageLen, or "Message not sent" otherwise.

// TIPS
// Here are some of the comparison operators in Go:

// == equal to
// != not equal to
// < less than
// > greater than
// <= less than or equal to
// >= greater than or equal to
package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
