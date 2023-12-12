// ASSIGNMENT
// Let's rewrite our printCleanReviews function from earlier using a for-of loop.

// The function should loop over the entire reviews array and print:

// Clean review: review

// Where review is the review text. But it should only do so if the review does not contain the given badWord.
const printCleanReviews = (reviews, badWord) => {
    // ?
    for (let review of reviews) {
        if (review.includes(badWord)) {
            continue
        }
        console.log(`Clean review: ${review}`)
    }
}

// don't touch below this line

printCleanReviews(['The movie sucks', 'I love it', 'What garbage', 'so sucky'], 'suck')
console.log('---')
printCleanReviews(['The movie sucks', 'I love it', 'What darn crap', 'darn, so sucky'], 'darn')
console.log('---')
