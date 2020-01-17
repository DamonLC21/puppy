package main

import "github.com/gin-gonic/gin"

// capital D means it's exported (public method)
// := is the go syntax for creating a new variable
// only inside a function
// go likes single letter variables
// H is a struct --  are objects/classes

func main() {
	router := gin.Default()
	router.GET("/puppy", handlePuppy)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
