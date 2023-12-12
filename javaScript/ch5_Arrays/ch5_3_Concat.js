// ASSIGNMENT
// Complete the uploadNewMovies function. It accepts the oldMovies array, which is the database of movies already in the MovieStarz system, and a newMovies array, which are the new movies we need to add to the system.

// It should return a new array containing all the movies from both arrays. First the old, then the new.
const uploadNewMovies = (oldMovies, newMovies) => {
    // ?
    return oldMovies.concat(newMovies)
}

// don't touch below this line

const oldMovies = ['Once Upon a Time', 'Django Unchained', 'The Hateful 8']
const newMovies = ['Inglorious Basterds', 'Dune']
const allMovies = uploadNewMovies(oldMovies, newMovies)

console.log('All movies database:')
logArray(allMovies)


function logArray(arr) {
    for (const a of arr) {
        console.log(` - ${a}`)
    }
    console.log('---')
}
