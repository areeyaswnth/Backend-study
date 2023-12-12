// ASSIGNMENT
// Fix the scoping bug and simplify the getIsPowerUser function. It should return true if the user has submitted more than 10 reviews to MovieStarz.

// Try to make it a 1-line function body if you can.
function getIsPowerUser(numReviews) {
    var isPowerUser = false
    if (numReviews > 10) {
        isPowerUser = true
    }
    return isPowerUser
}

// don't touch below this line

function test(numReviews) {
    const isPowerUser = getIsPowerUser(numReviews)
    console.log(`Number of reviews: ${numReviews}, is power user: ${isPowerUser}`)
}

test(100)
test(50)
test(10)
test(5)
test(2)
