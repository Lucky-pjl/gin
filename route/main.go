package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		// 重定向
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})

	r.GET("/a", func(c *gin.Context) {
		// 跳转到b对应的处理函数 转发
		c.Request.URL.Path="/b"
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"b"})
	})

	// 路由与路由组
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK,gin.H{"method":"GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK,gin.H{"method":"POST"})
		}
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"not found",
		})
	})

	// 路由组
	// 把公用的前缀提取，创建一个路由组
	vGroup := r.Group("/video")
	vGroup.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
	})

	r.Run(":9090")
}
