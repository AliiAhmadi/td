package todo

import "time"

// item struct represents todo items
type item struct {
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// List represents all todos in application
type List []item
