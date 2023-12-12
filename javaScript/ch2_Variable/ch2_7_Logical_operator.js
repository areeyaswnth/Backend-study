// ASSIGNMENT
// We need to be able to identify whether or not a movie is a "blockbuster" in MovieStarz.

// In order to be a blockbuster, a movie must fulfill the following criteria:

// It's a big hit
// It is new
// It has awards or it can have a sequel
// It isn't rated X
// Set the isBlockBuster using the given variables and the appropriate logical operators. Remember that you can specify an "order of operations" by using parentheses.
const isBigHit = true
const isNew = true
const hasAwards = false
const canHaveSequel = true
const isRatedX = false

// don't touch above this line

const isBlockBuster = isBigHit && isNew && (hasAwards || canHaveSequel) && !isRatedX

// don't touch below this line

console.log(`The movie is a blockbuster: ${isBlockBuster}`)
