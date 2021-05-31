package controller

import (
	"gin/demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateTodo(c *gin.Context) {
	// 1.取出数据
	var todo model.Todo
	c.BindJSON(&todo)
	// 2.存入数据库
	err := model.CreateTodo(&todo)
	// 3.返回响应
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	} else {
		c.JSON(http.StatusOK,todo)
	}
}

func GetTodos(c *gin.Context) {
	todos,err := model.GetTodos()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	} else {
		c.JSON(http.StatusOK,todos)
	}
}

func UpdateTodo(c *gin.Context) {
	id,_ := c.Params.Get("id")
	var todo model.Todo
	todo.ID,_ = strconv.Atoi(id)
	c.BindJSON(&todo)
	err := model.UpdateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	} else {
		c.JSON(http.StatusOK,todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id,_ := c.Params.Get("id")
	var todo model.Todo
	todo.ID,_ = strconv.Atoi(id)
	err := model.DeleteTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	} else {
		c.JSON(http.StatusOK,todo)
	}
}