package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index",func(c *gin.Context) {
		// 方法1.使用map
		//data := map[string]interface{}{
		//	"name":"zs",
		//	"message":"hello world!",
		//	"age":"20",
		//}
		data := gin.H{
			"name":"zs",
			"message":"hello world!",
			"age":"20",
		}
		c.JSON(http.StatusOK,data)
	})

	// 方法2.结构体
	type msg struct {
		Name string
		Message string
		Age int
	}
	r.GET("/struct",func(c *gin.Context) {
		data := msg{"zs", "Hello World!",30,
		}
		c.JSON(http.StatusOK,data)
	})
	r.Run(":9090")
}