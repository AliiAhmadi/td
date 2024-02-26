package todo

import (
	"fmt"
	"time"
)

// item struct represents todo items
type item struct {
	Task        string    `json:"task"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"updated_at"`
}

// List represents all todos in application
type List []item

// Add creates a new task and add that to list
func (l *List) Add(task string) {
	tk := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, tk)
}

// Complete makes a todo complete by change Done from false to true
func (l *List) Complete(index int) error {
	if index <= 0 || index > len(*l) {
		return fmt.Errorf("item %d does not exist", index)
	}

	(*l)[index-1].Done = true
	(*l)[index-1].CompletedAt = time.Now()
	return nil
}
