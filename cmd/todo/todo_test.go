package todo_test

import (
	"testing"

	"github.com/AliiAhmadi/td/cmd/todo"
)

// TestAdd tests add method on list by adding
// some tasks and check them
func TestAdd(t *testing.T) {
	l := todo.List{}

	names := []string{
		"task 1",
		"task 2",
		"task 3",
		"task 4",
		"task 5",
	}

	for i := 0; i < len(names); i++ {
		l.Add(names[i])

		if l[i].Task != names[i] {
			t.Errorf("TaskAdd: expected %s - got %s", names[i], l[i].Task)
		}
	}
}
