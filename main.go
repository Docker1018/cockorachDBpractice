package main

import (	
	"cockroach/controller"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	controller.SelectData()
	router := gin.Default()

	router.GET("/test/:name", getTest)

	router.Run(":8085")
}

func getTest(c *gin.Context) {
	data := controller.SelectData()
	c.JSON(200, data)
}
