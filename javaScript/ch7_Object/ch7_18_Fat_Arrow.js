// ASSIGNMENT
// The server logs that print when the MovieStarz program starts aren't working.

// Fix the bug.
const server = {
    port: 8080,
    name: 'MovieStarz',
    getLogs: function () {

        return `Starting ${this.name} server on ${this.port}`
    }
}

// don't touch below this line

const logs = server.getLogs()
console.log(logs)
