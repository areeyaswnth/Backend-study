// ASSIGNMENT
// We need to render all the "star rating" options to our MovieStarz users. They are allowed to give ratings between 1 and 10.

// Use the "decrement" operator (--) instead of the "increment" (++) operator to write a for-loop that prints:

// 10 stars
// 9 stars
// etc...
// 1 star
// Notice that 1 star is not plural.
for (let i = 10; i >= 1; i--) {
    if (i == 1) {
        console.log(`${i} star`)
        break
    }
    console.log(`${i} stars`)
}