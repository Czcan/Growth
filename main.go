package main
/*
controller主要放些调用业务逻辑的函数
dao层主要放与数据库相关的操作
这是小demo，所有model及model的增删改查操作直接放在models层。
大项目还有一个logic逻辑层，放调用modelsCRUD操作的业务逻辑函数，然后controller层的函数再调用logic的业务逻辑函数
router层放路由
static层放静态资源
templates层放模板文件
*/

import (
	"github.com/Czcan/bubble_demo/dao"
	"github.com/Czcan/bubble_demo/models"
	"github.com/Czcan/bubble_demo/router"
	"log"
)

func main(){
	//创建数据库
	//sql CREATE DATABASE bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil{
		log.Fatal(err)
	}
	defer dao.Close()  //程序退出关闭数据库连接
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := router.SetupRouter()
	r.Run()
}