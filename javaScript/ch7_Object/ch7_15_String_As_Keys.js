// ASSIGNMENT
// We need to keep track of the number of reviews we've captured for all the movies in our database. Complete the getCountsByTitle function. It should return an object where each key is the title of a movie, and the value is how many times that movie appeared in the movies parameter.

const getCountsByTitle = (movies) => {
    // ?
    let title = {

    }
    for (let key of movies) {
        if (!title[key]) {
            title[key] = 0
        }
        title[key]++
    }
    return title
}

// don't touch below this line

function test(movies) {
    const counts = getCountsByTitle(movies)
    for (const [movie, count] of Object.entries(counts)) {
        console.log(`'${movie}' has ${count} reviews`)
    }
    console.log('---')
}

test([
    'Ice Age',
    'The Forgotten',
    'The Northman',
    'American Psycho',
    'Ice Age',
    'Ice Age',
    'American Psycho'
])

test([
    'Big Daddy',
    'Big Daddy',
    'The Big Short',
    'The Big Short',
    'The Big Short'
])
