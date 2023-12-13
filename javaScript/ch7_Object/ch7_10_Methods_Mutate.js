// ASSIGNMENT
// Complete the submitReview() method on the author object. It should simply increment that author's numReviews counter by 1.
const author = {
    name: 'Simon Garfunkle',
    numReviews: 8,
    submitReview() {
        // ?
        this.numReviews++
    }
}

// don't touch below this line

console.log(`${author.name} has submitted ${author.numReviews} reviews`)

console.log(`${author.name} is submitting a review...`)
author.submitReview()
console.log(`${author.name} has submitted ${author.numReviews} reviews`)

console.log(`${author.name} is submitting a review...`)
author.submitReview()
console.log(`${author.name} has submitted ${author.numReviews} reviews`)

console.log(`${author.name} is submitting a review...`)
author.submitReview()
console.log(`${author.name} has submitted ${author.numReviews} reviews`)

console.log(`${author.name} is submitting a review...`)
author.submitReview()
console.log(`${author.name} has submitted ${author.numReviews} reviews`)
