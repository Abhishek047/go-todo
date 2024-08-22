package main

import (
	"fmt"

	apptype "github.com/Abhishek047/todo-app/app-type"
	"github.com/Abhishek047/todo-app/database"
)

func AddTodoHandler(text string) {
	DB := database.GetDb()
	todo := apptype.GetTodoList()
	todoItem := todo.AddTodo(text, DB)
	fmt.Println("Added new item")
	apptype.PrintTodoItems([]apptype.Todo{
		todoItem,
	})
}

func ListAllTodo() {
	DB := database.GetDb()
	todo := apptype.GetTodoList()
	todo.FetchList(DB)
	todo.ShowAll()
}

func MarkTodoDone(todoId string) {
	DB := database.GetDb()
	todo := apptype.GetTodoList()
	todo.MarkDone(todoId, DB)
}
