// ASSIGNMENT
// Keeping track of time in a message-sending application is critical. Imagine getting an appointment reminder an hour after your doctor's visit.

// Complete the code using a computed constant to print the number of seconds in an hour.
package main

import "fmt"

func ch2_9() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}
