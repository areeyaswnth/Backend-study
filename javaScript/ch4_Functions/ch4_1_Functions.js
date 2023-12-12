// ASSIGNMENT
// Complete the getLabel function. It takes the number of stars a movie received in its reviews and returns a string label describing how good the movie was.

// 8 - 10 stars: great
// 4 - 7 stars: okay
// less than or equal to 3 stars: awful
function getLabel(numStars) {
    // ?
    if (numStars >= 8 && numStars <= 10) {
        return "great"
    }
    else if (numStars >= 4 && numStars <= 7) {
        return "okay"
    } else {
        return "awful"
    }
}

// don't touch below this line

function test(numStars) {
    console.log(`Stars=${numStars}, The movie was ${getLabel(numStars)}`)
}

test(10)
test(9)
test(8)
test(7)
test(6)
test(5)
test(4)
test(3)
test(2)
test(1)
test(0)
