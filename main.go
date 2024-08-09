package main

import (
	"fmt"

	"github.com/Abhishek047/todo-app/database"
	"github.com/Abhishek047/todo-app/todo"
)

func main() {
	DB := database.GetDb()
	fmt.Println(DB.Fetch())
	// if err != nil {
	// 	fmt.Println("we have an error")
	// }
	todo := todo.GetTodoList()
	todo.AddTodo("First Text", nil)
	todo.AddTodo("Second Text", nil)
	todo.ShowAll()
}
