package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// 参数绑定
		var u User
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"error":err.Error(),
			})
		}
		c.JSON(http.StatusOK,gin.H{
			"status":"ok",
			"username":u.Username,
			"password":u.Password,
		})
	})
	r.Run(":9090")
}