// ASSIGNMENT
// We've been asked to "bucket" costs for an entire month into the cost occurring on each day of the month.

// Complete the getCostsByDay function. It should return a slice of float64s, where each element is the total cost for that day. The length of the slice should be equal to the number of days represented in the costs slice, including any days that have no costs, up to the last day represented in the slice.

// Days are numbered and start at 0.

// EXAMPLE
// Input in day, cost format:

// []cost{
//     {0, 4.0},
//     {1, 2.1},
//     {1, 3.1},
//     {5, 2.5},
// }
// Copy icon
// I would expect this result:

//	[]float64{
//	    4.0,
//	    5.2,
//	    0.0,
//	    0.0,
//	    0.0,
//	    2.5,
//	}
package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	// ?
	maxDay := 0
	for _, c := range costs {
		if c.day > maxDay {
			maxDay = c.day
		}

	}
	var s = make([]float64, maxDay+1)
	for _, c := range costs {
		s[c.day] += c.value
	}
	return s
}

// dont edit below this line

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}
