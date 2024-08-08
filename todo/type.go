package todo

import "time"

type Todo struct {
	Id        int64
	Text      string
	Status    bool
	EndDate   *time.Time
	CreatedAt *time.Time
}

type TodoList struct {
	List []Todo
}
