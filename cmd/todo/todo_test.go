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

// TestComplete creates a list and add a task to that
// after that check for Done in that
func TestComplete(t *testing.T) {
	l := todo.List{}

	task := "here is a task for testing"
	l.Add(task)

	if l[0].Task != task {
		t.Errorf("TestComplete: expected %s - got %s", task, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("TaskComplete: new task should not be completed")
	}

	l.Complete(1)
	if !l[0].Done {
		t.Errorf("TaskComplete: task should be completed")
	}
}

// TestDelete tests delete method on list
func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"first task",
		"second task",
		"third task",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("TestDelete: expected %s - got %s", tasks[0], l[0].Task)
	}

	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("TestDelete: expected %d - got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("TestDelete: expected %s - got %s", tasks[2], l[1].Task)
	}
}
