// ASSIGNMENT
// Complete the movieExists function. It should loop over the entire movies array and print:

// Looking at: movie

// Where movie is the movie title in the movies array.

// Once it finds the movie title that matches the given movie argument, it should also print:

// Found: movie

// Where movie is the movie title.

// Then it should break from the loop and not print anything else.
const movieExists = (movies, movie) => {
    // ?
    for (let i = 0; i < movies.length; i++) {
        console.log('Looking at: '.concat(movies[i]))
        if (movies[i] == movie) {
            console.log('Found: '.concat(movie))
            break
        }

    }

}

// don't touch below this line

movieExists(['Incredibles', 'Tangled', 'Frozen'], 'Frozen')
console.log('---')
movieExists(['The Quick and the Dead', 'The Magnificent 7', 'Frozen'], 'The Magnificent 7')
console.log('---')
movieExists(['Dead', 'Alive', 'Flight', 'Rocky'], 'Flight')
console.log('---')
movieExists(['Dead', 'Alive', 'Flight', 'Rocky'], 'None')
console.log('---')
