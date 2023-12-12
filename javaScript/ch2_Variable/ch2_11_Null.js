// ASSIGNMENT
// We need to identify which movies in our database were submitted without titles! Unfortunately, we're storing the movie data incorrectly.

// Fix the code so that instead of storing a movie without a title as the string 'null', we use a null value.
const movieTitle = null

// don't touch below this line

// The 'not' operator implicitly
// casts movieTitle to a boolean value
const movieHasNoTitle = !movieTitle

console.log(`The movie does not have a title: ${movieHasNoTitle}`)
