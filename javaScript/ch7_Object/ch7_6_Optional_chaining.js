// ASSIGNMENT
// Our MovieStarz database is pretty smart - it will only save a value if the value is "truthy" in JavaScript. If it's "falsy", it will just ignore it without errors.

// Complete the getState function. It takes a review as input and returns the state string if the review has one. Otherwise, it returns undefined (a falsy value).

// You should be able to do this in 1 line with the optional chaining operator. Take a look at the test cases to see the shape of the reviews objects.
function getState(review) {
    // ?
    return review.author?.location?.state
}

// don't touch below this line

function test(review) {
    const state = getState(review)
    if (state) {
        console.log(`Adding ${state} to the database...`)
    } else {
        console.log('No state found...')
    }
}

test({
    text: 'This movie was awful',
    stars: 2,
    author: {
        firstName: 'Johnny',
        lastName: 'Comelately',
        createdAt: '2022-08-17T15:41:25+00:00',
        location: {
            state: 'Utah'
        }
    }
})

test({
    text: 'This movie was okay...',
    stars: 5
})

test({
    text: 'This movie was awful',
    stars: 2,
    author: {
        firstName: 'Jill',
        lastName: 'Comelately',
        createdAt: '2022-08-17T15:41:25+00:00',
        location: {
            state: 'Nevada'
        }
    }
})

test({
    text: 'This movie was awful',
    stars: 2,
    author: {
        firstName: 'George',
        lastName: 'Jimenez',
        createdAt: '2022-08-17T15:41:25+00:00'
    }
})
