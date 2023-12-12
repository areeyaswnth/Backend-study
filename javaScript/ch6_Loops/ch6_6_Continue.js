// ASSIGNMENT
// Complete the printCleanReviews function. It should loop over the entire reviews array and print:

// Clean review: review

// Where review is the review text. But it should only do so if the review does not contain the given badWord.

// Do not use an else-statement!!! Use the continue keyword instead to accomplish the same thing.
const printCleanReviews = (reviews, badWord) => {
    // ?
    for (let i = 0; i < reviews.length; i++) {
        if (reviews[i].includes(badWord)) {
            continue
        }
        console.log(`Clean review: ${reviews[i]}`)
    }
}

// don't touch below this line

printCleanReviews(['The movie sucks', 'I love it', 'What garbage', 'so sucky'], 'suck')
console.log('---')
printCleanReviews(['The movie sucks', 'I love it', 'What darn crap', 'darn, so sucky'], 'darn')
console.log('---')
