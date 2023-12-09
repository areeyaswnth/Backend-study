// ASSIGNMENT
// Complete the maxMessages function. Given a cost threshold, it should calculate the maximum number of messages that can be sent.

// Each message costs 1.0, plus an additional fee. The fee structure is:

// 1st message: 1.0 + 0.00
// 2nd message: 1.0 + 0.01
// 3rd message: 1.0 + 0.02
// 4th message: 1.0 + 0.03
// BROWSER FREEZE
// If you lock up your browser by creating an infinite loop that isn't breaking, just click the cancel button.
package main

import (
	"fmt"
)

func maxMessages(thresh float64) int {
	// ?
	total := 0.00
	num := 0
	for {

		total += 1.0 + (float64(num) * 0.01)

		if total > thresh {
			break
		}
		num++
	}
	return num
}

// don't edit below this line

func test(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(5.1)
	test(40.00)
	test(50.00)
}
