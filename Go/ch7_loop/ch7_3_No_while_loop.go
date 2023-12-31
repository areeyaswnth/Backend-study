// ASSIGNMENT
// We have an interesting new cost structure from our SMS vendor. They charge exponentially more money for each consecutive text we send! Let's write a function that can calculate how many messages we can send in a given batch given a costMultiplier and a maxCostInPennies.

// In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied by the costMultiplier. That gets expensive!

// There is an infinite loop in the code! Add a condition to the for loop to fix the bug. The loop should break as soon as the actualCostInPennies is greater than the maxCostInPennies.
package main

import (
	"fmt"
)

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for {
		actualCostInPennies *= costMultiplier
		if actualCostInPennies > float64(maxCostInPennies) {
			break
		}
		maxMessagesToSend++

	}

	return maxMessagesToSend
}

// don't touch below this line

func test(costMultiplier float64, maxCostInPennies int) {
	maxMessagesToSend := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Multiplier: %v\n",
		costMultiplier,
	)
	fmt.Printf("Max cost: %v\n",
		maxCostInPennies,
	)
	fmt.Printf("Max messages you can send: %v\n",
		maxMessagesToSend,
	)
	fmt.Println("====================================")
}

func main() {
	test(1.1, 5)
	test(1.3, 10)
	test(1.35, 25)
}
