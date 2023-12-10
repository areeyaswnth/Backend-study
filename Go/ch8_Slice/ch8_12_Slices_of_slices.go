// ASSIGNMENT
// We support various graphs and dashboards on Textio that display message analytics to our users. The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.

// Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers where the value of each cell is i * j where i and j are the indexes of the row and column respectively.

// For example, a 10x10 matrix would look like this:

// [0 0 0 0 0 0 0 0 0 0]
// [0 1 2 3 4 5 6 7 8 9]
// [0 2 4 6 8 10 12 14 16 18]
// [0 3 6 9 12 15 18 21 24 27]
// [0 4 8 12 16 20 24 28 32 36]
// [0 5 10 15 20 25 30 35 40 45]
// [0 6 12 18 24 30 36 42 48 54]
// [0 7 14 21 28 35 42 49 56 63]
// [0 8 16 24 32 40 48 56 64 72]
// [0 9 18 27 36 45 54 63 72 81]
package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	// ?
	matrix := [][]int{}
	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)

	}
	return matrix
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
