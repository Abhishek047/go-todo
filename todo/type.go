package todo

import "time"

type Todo struct {
	text      string
	status    bool
	endDate   *time.Time
	createdAt *time.Time
}

type ListTodo struct {
	todoList []Todo
}
