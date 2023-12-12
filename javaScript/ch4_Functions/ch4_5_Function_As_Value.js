// ASSIGNMENT
// After submitting a review to MovieStarz, we need to print the review itself to our audit logs, then save it to the database!

// Complete the onReview and main functions.

// ONREVIEW()
// This is the function that's called when a new review is submitted by a user. It should log the review text and then call the provided callbackFunction and give it the reviewText as input.

// MAIN()
// Call onReview with the ohBrotherWhereArtThouReview review, then with the twentyTwelveIceAgeReview review. In both cases use the saveToDatabase function as the callback.
function onReview(reviewText, callbackFunction) {
    console.log(`Review: ${reviewText}`);
    callbackFunction(re)
    // ?

}

function main() {
    const ohBrotherWhereArtThouReview = 'Stellar movie!'
    const twentyTwelveIceAgeReview = 'Not my favorite'
    // ?
    onReview(ohBrotherWhereArtThouReview, saveToDatabase)
    onReview(twentyTwelveIceAgeReview, saveToDatabase)

}

// Don't edit below this line

function saveToDatabase(reviewText) {
    console.log(`Saving '${reviewText}' to database...`)
}

main()
