package controller
import (
	"github.com/Czcan/bubble_demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
/*
	  url   --> controller -->  logic    -->  models
	请求来了 -->    控制器   -->  业务逻辑  -->  模型层的增删改查
*/
func IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",nil)
}

func CreateATodo(c *gin.Context) {
	//前端页面填写待办事项，点击提交，会发送请求到这里
	//1.从请求中拿数据
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	//err := DB.Create(&todo).Error
	//if err != nil{
	//}
	//3.返回响应
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK,gin.H{
		//	"code" : 2000,
		//	"msg" : "success",
		//	"data" : todo,
		//})
	}
}

func GetTodoList(c *gin.Context){

		if todoList, err := models.GetAllTodo(); err != nil{
			c.JSON(http.StatusOK, gin.H{"error":err.Error()})
		}else {
			c.JSON(http.StatusOK, todoList)
		}
}

func UpdateATodo(c *gin.Context){
	id, ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
	if err := models.DeleteATodo(id); err != nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id:"had been deleted"})
	}
}


