package routers

import (
	"gin/demo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo",controller.CreateTodo)
		// 查看所有代办事项
		v1Group.GET("/todo",controller.GetTodos)
		// 修改
		v1Group.PUT("/todo/:id",controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id",controller.DeleteTodo)
	}
	return r
}