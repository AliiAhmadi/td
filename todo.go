package td

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item struct represents todo items
type item struct {
	Task        string    `json:"task"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
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
	err := l.indexCheck(index)
	if err != nil {
		return err
	}

	(*l)[index-1].Done = true
	(*l)[index-1].CompletedAt = time.Now()
	return nil
}

// Uncomplete a completed todo
func (l *List) Uncomplete(index int) error {
	if err := l.indexCheck(index); err != nil {
		return err
	}

	(*l)[index-1].Done = false
	(*l)[index-1].CompletedAt = time.Time{}
	return nil
}

// Delete method deletes a todo from list by its index
func (l *List) Delete(index int) error {
	err := l.indexCheck(index)
	if err != nil {
		return err
	}

	*l = append((*l)[:index-1], (*l)[index:]...)
	return nil
}

// Save converts list of items to a json and write them
// in input file
func (l *List) Save(file string) error {
	j, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(file, j, 0644)
}

// Get reads specified file and unmarshal its to
// a list of items
func (l *List) Get(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return nil
		default:
			return err
		}
	}

	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, l)
}

// indexCheck checks index to be in correct context
func (l *List) indexCheck(index int) error {
	if index <= 0 || index > len(*l) {
		return fmt.Errorf("item %d does not exist", index)
	}

	return nil
}

// String returns formatted list
func (l *List) String() string {
	formatted := ""

	for k, v := range *l {
		prefix := "   "

		if v.Done {
			prefix = "X  "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, v.Task)
	}

	return formatted
}

func (l *List) Format() string {
	formatted := ""

	for k, v := range *l {
		prefix := "   "

		if v.Done {
			continue
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, v.Task)
	}

	return formatted
}
