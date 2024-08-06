package todo

import "time"

func CreateTodoItem(text string, endDate *time.Time) Todo {
	now := time.Now()
	return Todo{
		Text:      text,
		Status:    false,
		EndDate:   endDate,
		CreatedAt: &now,
	}
}
