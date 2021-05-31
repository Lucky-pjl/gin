package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()

	// 加载静态文件
	r.Static("/xxx","./statics")

	// gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe" : func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("templates/index.tmpl") // 解析模板
	r.GET("/index",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{ // 渲染模板
			"title":"<a href='https://baidu.com'>百度</a>",
		})
	})

	r.Run(":9090")
}