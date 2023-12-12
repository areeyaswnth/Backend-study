// ASSIGNMENT
// First, print a slice of the movies array that starts with the third item in the array and goes to the end.
// Next, print a slice of the movies array that ends with the third to last item in the array and starts at the beginning.
// Use logArray() for printing the arrays.

// EXAMPLE OF POSITIONING
// // 2 is the third item
// // 3 is the "third to last" item
// [0, 1, 2, 3, 4, 5]
const movies = [
    'oh brother where art thou',
    'oceans eleven',
    'fight club',
    'the island',
    'shutter island',
    'the magnificent seven'
]

function logArray(arr) {
    for (const a of arr) {
        console.log(` - ${a}`)
    }
    console.log('---')
}

// don't touch above this line

// ?
logArray(movies.slice(2))
logArray(movies.slice(0, -2))
