package main

import "github.com/gin-gonic/gin"


var Db []Recipe = initDB()

func main() {
	r := gin.Default()

	routes(r)

	r.Run()
}
