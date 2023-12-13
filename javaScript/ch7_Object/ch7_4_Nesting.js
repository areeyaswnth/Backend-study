// ASSIGNMENT
// In most systems, "entities" are stored as objects. For example, "users", "reviews", and "actors" might all be different "entities" within the MovieStarz system.

// We store movie reviews as nested JavaScript objects! Print the given review's author's first name by using the . operator.



const review = {
    text: 'This movie was awful',
    stars: 2,
    author: {
        firstName: 'Johnny',
        lastName: 'Comelately',
        createdAt: '2022-08-17T15:41:25+00:00'
    }
}

// don't touch above this line

// ?

console.log(review.author.firstName)