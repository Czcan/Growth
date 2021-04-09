package router

import (
	"github.com/Czcan/bubble_demo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	//告诉gin框架去哪里找模板文件引用的静态文件
	r.Static("/static","static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/bubble", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}