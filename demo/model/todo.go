package model

import "gin/demo/dao"

// Todo Model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

// Todo crud
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return err
}

func GetTodos() (todos []*Todo,err error) {
	err = dao.DB.Find(&todos).Error
	return todos,err
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Model(todo).Update("status",todo.Status).Error
	return err
}

func DeleteTodo(todo *Todo) (err error) {
	err = dao.DB.Delete(todo).Error
	return err
}