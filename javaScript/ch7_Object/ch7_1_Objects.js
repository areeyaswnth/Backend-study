// ASSIGNMENT
// Complete the getMovieRecord function. It takes a movie's title, the number of stars, and the username of the person who gave the star rating. It should return an object with 4 properties.

// title
// stars
// username
// id
// Where the id is the title and the username concatenated together with a hyphen for uniqueness. For example, toy story and hanks123 would make the id toy story-hanks123.
function getMovieRecord(title, stars, username) {
    // ?
    const id = title + '-' + username
    const object = {
        title: title,
        stars: stars,
        username: username,
        id: id
    }
    return object
}

// Don't touch below this line

logObject(getMovieRecord('oh brother where art thou', 3, 'wagslane'))
logObject(getMovieRecord('frozen', 5.5, 'elonmusk'))
logObject(getMovieRecord('toy story', 4, 'prince'))

function logObject(obj) {
    for (const key in obj) {
        console.log(` - ${key}: ${obj[key]}`)
    }
    console.log('---')
}
