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

func PrintTodoItems(todoItem []Todo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTodo\tStatus\tCreatedAt")
	for _, item := range todoItem {
		printStr := fmt.Sprintf("%v\t%v\t%v\t%v", item.Id, item.Text, item.Status, item.CreatedAt.Format("Jan 02, 2006"))
		fmt.Fprintln(w, printStr)
	}
	w.Flush()
}

func (todo *TodoList) AddTodo(text string, dbInstance DbI) Todo {
	todoItem := CreateTodoItem(text, nil)
	dbInstance.AddItem(todoItem)
	return todoItem
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
	PrintTodoItems(todo.List)
}

func (todo *TodoList) FetchList(dbInstance DbI) error {
	if dbInstance != nil {
		newList, err := dbInstance.Fetch()
		if err != nil {
			return err
		}
		todo.List = newList
	}
	return nil
}

func (todo *TodoList) MarkDone(id string, dbInstance DbI) error {
	if dbInstance != nil {
		updatedItem := dbInstance.UpdateTodoDone(id)
		if updatedItem != nil {
			PrintTodoItems([]Todo{
				*updatedItem,
			})
		}
	} else {
		fmt.Println("No DB to work with")
	}
	return nil
}

// can make it a singleton so only to work on 1 todo list
func GetTodoList() *TodoList {
	todoList := TodoList{
		List: []Todo{},
	}
	return &todoList
}
