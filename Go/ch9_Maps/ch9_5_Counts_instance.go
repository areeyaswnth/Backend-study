// ASSIGNMENT
// We have a slice of user ids, and each instance of an id in the slice indicates that a message was sent to that user. We need to count up how many times each user's id appears in the slice to track how many messages they received.

// Implement the getCounts function. It should return a map of string -> int so that each int is a count of how many times each string was found in the slice.
package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func getCounts(userIDs []string) map[string]int {
	// ?
	users := map[string]int{}
	for _, user := range userIDs {
		users[user]++
	}
	return users
}

// don't edit below this line

func test(userIDs []string, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)
	}
	fmt.Println("=====================================")
}

func main() {
	userIDs := []string{}
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:2])
	}

	test(userIDs, []string{"00", "ff", "dd"})
	test(userIDs, []string{"aa", "12", "32"})
	test(userIDs, []string{"bb", "33"})
}
