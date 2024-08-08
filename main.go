package main

import (
	"github.com/Abhishek047/todo-app/todo"
)

func main() {
	// data, err := configs.GetConfigData()
	// if err != nil {
	// 	fmt.Println("we have an error")
	// }
	todo := todo.GetTodoList()
	todo.AddTodo("First Text", nil)
	todo.AddTodo("Second Text", nil)
	todo.ShowAll()
}
