package main

import (
	"gin/demo/dao"
	"gin/demo/model"
	"gin/demo/routers"
)

func main() {

	// 连接数据库
	defer dao.DB.Close()
	dao.DB.AutoMigrate(&model.Todo{})

	r := routers.SetupRouter()
	r.Run(":9000")
}
