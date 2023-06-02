var golang = require('golang')
    golang.build(`
        package main
        import "fmt"
        func main() {
            fmt.Println("hello world")
        }
    `).then((data) => {
        // output on terminal + directory of build
		console.log(data)
    }).catch((e) => {
        throw e // err code.
    })