package todo

import "time"

func CreateTodoItem(text string, endDate *time.Time) Todo {
	now := time.Now()
	return Todo{
		text:      text,
		status:    false,
		endDate:   endDate,
		createdAt: &now,
	}
}
