// ASSIGNMENT
// At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send.

// COMPLETE THE BULKSEND() FUNCTION
// This function should return the total cost (as a float64) to send a batch of numMessages messages. Each message costs 1.0, plus an additional fee. The fee structure is:

// 1st message: 1.0 + 0.00
// 2nd message: 1.0 + 0.01
// 3rd message: 1.0 + 0.02
// 4th message: 1.0 + 0.03
// ...
// Use a loop to calculate the total cost and return it.
package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	// ?

	total := 0.00
	for i := 0; i < numMessages; i++ {
		fee := float64(i) * 0.01
		total += 1.0 + fee
	}
	return total
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}
