package main

import (
	"fmt"

	apptype "github.com/Abhishek047/todo-app/app-type"
	"github.com/Abhishek047/todo-app/database"
)

func main() {
	DB := database.GetDb()
	todo, err := apptype.GetTodoList(DB)
	if err != nil {
		fmt.Println("fetch errror")
	}
	todo.AddTodo("New Text", nil, DB)
	// todo.AddTodo("Second LDO Text", nil, DB)
	todo.ShowAll()
}
