// ASSIGNMENT
// When searching through MovieStarz codebase, you found a bug dealing with scope issues!

// isMovieValid should return true if the movie title's length is greater than 2. Use the movieLength function to get the length of a movie title and fix the scoping bug.
function isMovieValid(title) {
    function movieLength(title) {
        const len = title.length
        return len
    }

    // don't touch above this line

    const len = movieLength(title)
    return len > 2
}

// don't touch below this line

function test(title) {
    const valid = isMovieValid(title)
    console.log(`'${title}' is valid: ${valid}`)
}

test('The League of Extraordinary Gentlemen')
test('Hint for Red October')
test('007')
test('')
test('12')
