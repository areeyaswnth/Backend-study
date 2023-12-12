// ASSIGNMENT
// One of the senior developers on the MovieStarz team wrote a function called isClean that detects whether or not a review has any bad words in it. It makes use of JavaScript's built-in .includes() method.

// Fix the bug in the function! It should return false if dang, shoot, or heck are present in the text, and true otherwise.
function isClean(review) {
    let clean = true
    if (review.includes('dang')) {
        clean = false
    }
    if (review.includes('shoot')) {
        clean = false
    }
    if (review.includes('heck')) {
        clean = false
    }
    return clean
}


// Don't edit below this line

function test(review) {
    const clean = isClean(review)
    console.log(`'${review}' is clean: ${clean}`)
}

test('avril lavigne has best dang tour')
test('what a bad film')
test('oh my heck I hated it')
test('ripoff')
test('That was a pleasure')
test('shoot! I cant say I liked it')
