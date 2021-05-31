package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})

	// /login post 获取表单参数
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name" : username,
			"Password":password,
		})
	})

	// 获取路径参数
	r.GET("/path/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name" : name,
			"Age":age,
		})
	})
	r.Run(":9090")
}