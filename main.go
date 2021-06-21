package main

import (	
	"cockroach/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	controller.SelectData()
	// db.Action()
	// fmt.Println("?");
	router := gin.Default()

	router.GET("/test/:name", getTest)

	router.Run(":8085")
}

func getTest(c *gin.Context) {
	name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
}
