// ASSIGNMENT
// Let's fix up our profanity detection to make it a little more robust. Rather than just marking a review as "clean" or "not clean" we need to give it a ranking, which we'll represent as one of 3 strings:

// clean: No bad words
// dirty: 1 bad word
// filthy: 2 or more different bad words
// The bad words are:

// dang
// shoot
// heck
// Complete the getCleanRank function. If a word contains punctuation, it should escape detection in our naive algorithm for now.
function getCleanRank(reviewWords) {
    const dang = reviewWords.includes('dang') ? 1 : 0;
    const shoot = reviewWords.includes('shoot') ? 1 : 0;
    const heck = reviewWords.includes('heck') ? 1 : 0;
    const sum = dang + shoot + heck;

    if (sum === 0) {
        return "clean";
    } else if (sum === 1) {
        return "dirty";
    } else {
        return "filthy";
    }
}



// Don't edit below this line

function test(reviewWords) {
    const cleanRank = getCleanRank(reviewWords)
    console.log(`'${reviewWords}' has rank: ${cleanRank}`)
}

test(['avril', 'lavigne', 'has', 'best', 'dang', 'tour'])
test(['what', 'a', 'bad', 'film'])
test(['oh', 'my', 'heck', 'I', 'hated', 'it'])
test(['ripoff'])
test(['That', 'was', 'a', 'pleasure'])
test(['shoot!', 'I', 'cant', 'say', 'I', 'liked', 'the', 'dang', 'thing'])
test(['shoot', 'dang', 'heck'])
