// ASSIGNMENT
// As it turns out, we have a bunch of old movie records in our MovieStarz database that never had an ID assigned in the first place! Write the addID function. It takes an existing movie record object and calculates the id:

// id = title-username

// Where title is the movie record's title property and username is the movie records username property.

// Add the id to the id property of the movie record and return the new movie record object.



// // 
function addID(movieRecord) {
    // ?
    movieRecord.id = movieRecord.title + "-" + movieRecord.username
    return movieRecord

}

// Don't touch below this line

function getMovieRecord(title, stars, username) {
    return {
        title: title,
        stars: stars,
        username: username
    }
}

function logObject(obj) {
    for (const key in obj) {
        console.log(` - ${key}: ${obj[key]}`)
    }
    console.log('---')
}

const record = getMovieRecord('Frozen', 8.5, 'Elsa')
logObject(record)
console.log('Adding ID...')
const idRecord = addID(record)
logObject(idRecord)
