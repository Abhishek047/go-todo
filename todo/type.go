package todo

import "time"

type Todo struct {
	Text      string
	Status    bool
	EndDate   *time.Time
	CreatedAt *time.Time
}

type ListTodo struct {
	TodoList []Todo
}
