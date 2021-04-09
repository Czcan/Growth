package models

import "github.com/Czcan/bubble_demo/dao"

//Todo Model
type Todo struct {
	ID int  `json:"id"`
	Title string `json:"title"` //待办事项名字
	Status bool `json:"status"` //状态：已完成/未完成
}

/*
	Todo这个Model的增删改查操作都在这里
*/

//创建todo
func CreateATodo(todo *Todo) error{
	err := dao.DB.Create(&todo).Error
	return err
}

//查看所有todo
func GetAllTodo()(todoList []*Todo, err error){
	err = dao.DB.Find(&todoList).Error
	if err != nil{
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error){
	todo = new(Todo)
	err = dao.DB.Where("id=?",id).First(&todo).Error
	if err != nil{
		return  nil , err
	}
	return
}


func UpdateATodo(todo *Todo)error{
	err := dao.DB.Save(&todo).Error
	return err
}

func DeleteATodo(id string)error{
	err := dao.DB.Where("id=?",id).Delete(Todo{}).Error  //为什么Delete()传个空struct
	return err
}

