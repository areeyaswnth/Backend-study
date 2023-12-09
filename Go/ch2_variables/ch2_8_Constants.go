//ASSIGNMENT
// USE TWO SEPARATE CONSTANTS
// Something weird is happening in this code.

// What should be happening is that we create 2 separate constants: premiumPlanName and basicPlanName. Right now it looks like we're trying to overwrite one of them.
package main

import "fmt"

func ch2_8() {
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
