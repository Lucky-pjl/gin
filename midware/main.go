package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index in ...")
	c.JSON(http.StatusOK,gin.H{
		"msg":"index",
	})
}

// 定义一个中间件m1
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的函数
}

func authMiddleWare() gin.HandlerFunc {
	// 连接数据库
	// 或者做其他准备工作
	return func(c *gin.Context) {
		// 具体的逻辑
		// 是否登录判断
	}
}

// 中间件
func main() {
	r := gin.Default()
	// 全局注册中间件函数
	r.Use(m1)
	r.GET("/m1",indexHandler)
	r.GET("/m2", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"msg":"m2"})
	})
	r.Run(":9090")
}
