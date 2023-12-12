// ASSIGNMENT
// Complete the getMostRecentUser function. If the usernames array is empty, it should return null. Otherwise, it should return the last username in the array.

// The goal here is to find the user who signed up to MovieStarz most recently!
const getMostRecentUser = (usernames) => {
    // ?
    if (usernames.length <= 1) {
        return null
    }
    else {
        return usernames[usernames.length - 1]
    }

}

// don't touch below this line

console.log(`Most recent user: ${getMostRecentUser(
    []
)}`)

console.log(`Most recent user: ${getMostRecentUser(
    [
        'johndoe123',
        'billyrae456'
    ]
)}`)

console.log(`Most recent user: ${getMostRecentUser(
    [
        'wagslane',
        'jimmyjohn',
        'bopeep',
        'strightkilla',
        'reddyman'
    ]
)}`)

