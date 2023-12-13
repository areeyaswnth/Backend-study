// ASSIGNMENT
// Complete the getFirstReview() method. The method should return the first review in the user's review array.
const user = {
    getFirstReview() {
        // ?
        return this.reviews[0]
    },
    reviews: ['I hate Ice Age', 'I didn\'t enjoy it at all', 'What a fabulous film'],
    name: 'Bob Doogle'
}

// don't touch below this line

console.log(`${user.name}'s first review is: ${user.getFirstReview()}`)