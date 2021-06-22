package main

import (
	"cockroach/controller"
	"cockroach/model"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	controller.SelectData()
	router := gin.Default()

	// 會員名單
	router.GET("/member/list", getMemberList)

	// 註冊會員
	router.POST("/member/add", registerMember)

	// 更新會員
	router.PUT("/member/:account", updateMember)

	// 刪除會員
	router.DELETE("/member/:account", deleteMember)

	router.Run(":8085")
}

func getMemberList(c *gin.Context) {
	c.JSON(200, controller.SelectData())
}

func registerMember(c *gin.Context) {
	params := &model.Member{}
	c.BindJSON(params)
	c.JSON(200, controller.InsertData(params))
}

func updateMember(c *gin.Context) {
	acc := c.Param("account")
	c.JSON(200, acc)
}

func deleteMember(c *gin.Context) {
	acc := c.Param("account")
	c.JSON(200, acc)
}
