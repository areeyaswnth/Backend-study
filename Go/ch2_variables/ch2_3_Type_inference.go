// ASSIGNMENT
// Our current price to send a text message is 2 cents. However, it's likely that in the future the price will be a fraction of a penny, so we should use a float64 to store this value.

// Edit the penniesPerText declaration so that it's inferred by the compiler to be a float64.
package main

import "fmt"

func ch2_3() {
	penniesPerText := 2.22
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}
