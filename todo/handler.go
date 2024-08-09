package todo

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func CreateTodoItem(text string, endDate *time.Time) Todo {
	now := time.Now()
	currentTimestamp := now.UnixNano() / int64(time.Microsecond)
	return Todo{
		Id:        currentTimestamp,
		Text:      text,
		Status:    false,
		EndDate:   endDate,
		CreatedAt: &now,
	}
}

func (todo *TodoList) AddTodo(text string, endDate *time.Time) {
	todoItem := CreateTodoItem(text, endDate)
	todo.List = append(todo.List, todoItem)
	// do the save strategy here
}

func (todo *TodoList) DeleteTodo(id int64) {
	var indexToDelete int = -1
	for index, item := range todo.List {
		if item.Id == id {
			indexToDelete = index
			break
		}
	}
	if indexToDelete == -1 {
		fmt.Println("Id not found")
	} else {
		todo.List = append(todo.List[:indexToDelete], todo.List[indexToDelete+1])
	}
	// do the save strategy here
}

func (todo *TodoList) ShowAll() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTodo\tStatus\tCreatedAt")
	for _, item := range todo.List {
		printStr := fmt.Sprintf("%v\t%v\t%v\t%v", item.Id, item.Text, item.Status, item.CreatedAt.Format("Jan 02, 2006"))
		fmt.Fprintln(w, printStr)
	}
	w.Flush()
}

// can make it a singleton so only to work on 1 todo list
func GetTodoList() TodoList {
	todoList := TodoList{
		List: []Todo{},
	}
	// fetch here
	return todoList
}
