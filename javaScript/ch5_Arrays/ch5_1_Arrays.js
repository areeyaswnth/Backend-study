// ASSIGNMENT
// Create an empty array called 'movies'
// Push the dark knight to your array of movies
// Print the array using logArray()
// Push the notebook to your array of movies
// Print the array using logArray()
// Print the item at index 0 of your movies array
// ?
const movies = []
movies.push("the dark knight")
logArray(movies)
movies.push("the notebook")
logArray(movies)
console.log(movies[0])
// don't touch below this line

function logArray(arr) {
    console.log('logging array...')
    for (const a of arr) {
        console.log(` - ${a}`)
    }
    console.log('---')
}
