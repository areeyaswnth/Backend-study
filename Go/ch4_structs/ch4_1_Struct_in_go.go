// ASSIGNMENT
// Complete the messageToSend struct definition. It needs two fields:

// phoneNumber - an integer
// message - a string.
package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

// don't edit below this line

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func ch4_1() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
}
