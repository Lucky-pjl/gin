package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器请求携带的query string参数
		name := c.Query("name")
		//name := c.DefaultQuery("name","默认值")
		// name,ok := c.GetQuery("name")
		c.JSON(http.StatusOK,"hello " + name)
	})
	r.Run(":9090")
}