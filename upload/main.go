package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		file,err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		} else {
			dst := fmt.Sprintf("./%s",file.Filename)
			c.SaveUploadedFile(file,dst)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})
	r.Run(":9090")
}
