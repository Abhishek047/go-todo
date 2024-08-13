package apptype

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

func (todo *TodoList) AddTodo(text string, endDate *time.Time, dbInstance DbI) {
	todoItem := CreateTodoItem(text, endDate)
	todo.List = append(todo.List, todoItem)
	// do the save strategy here
	dbInstance.Save(todoItem)
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
func GetTodoList(dbInstance DbI) (TodoList, error) {
	todoList := TodoList{
		List: []Todo{},
	}
	if dbInstance != nil {
		newList, err := dbInstance.Fetch()
		if err != nil {
			return todoList, err
		}
		todoList.List = newList
	}
	// fetch here
	return todoList, nil
}
