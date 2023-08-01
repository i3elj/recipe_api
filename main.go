package main

import "github.com/gin-gonic/gin"


var Db []Recipe = initDB()

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	routes(r)

	r.Run()
}
