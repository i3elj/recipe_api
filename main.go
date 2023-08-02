package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	routes(r)

	r.Run()
}
