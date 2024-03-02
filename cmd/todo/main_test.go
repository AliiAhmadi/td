package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "td_test"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	bc := exec.Command("go", "build", "-o", binName)

	if err := bc.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build %s: %s", binName, err)
		os.Exit(1)
	}

	res := m.Run()
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(res)
}

func TestListTodosCLI(t *testing.T) {
	base := []string{
		"-task",
	}
	task := "test-task-(1)"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(dir, binName)

	t.Run("adding a todo from command line", func(t *testing.T) {
		cmd := exec.Command(path, append(base, task)...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	// Check
	t.Run("listing todos", func(t *testing.T) {
		cmd := exec.Command(path, "-all")

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := task + "\n"

		if string(out) != expected {
			t.Errorf("expected %q - got %q", expected, string(out))
		}
	})
}

func TestShowAllTasks(t *testing.T) {
	// base := []string{
	// 	"-all",
	// }

	// tasks := []string{
	// 	"here-is-first-task",
	// 	"here-is-second-task",
	// }

	// dir, err := os.Getwd()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// path := filepath.Join(dir, binName)

	// t.Run("adding todos from command line by using executable", func(t *testing.T) {
	// 	for _, task := range tasks {
	// 		cmd := exec.Command(path, append(base, task)...)
	// 		if err := cmd.Run(); err != nil {
	// 			t.Fatal(err)
	// 		}
	// 	}
	// })
}
