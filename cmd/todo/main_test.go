package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
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

func TestTodoCLI(t *testing.T) {
	base := []string{
		"-task",
	}
	tasks := []string{
		"test-task-(1)",
		"test-task-(2)",
		"test-task-(3)",
		"test-task-(4)",
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(dir, binName)

	// Completing all tasks and check -all and -list
	t.Run("completing all tasks and check -all and -list flags", func(t *testing.T) {

		for _, task := range tasks {
			cmd := exec.Command(path, append(base, task)...)
			if err := cmd.Run(); err != nil {
				t.Fatal(err)
			}
		}

		allcmd := exec.Command(path, "-all")

		out, err := allcmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := stringer(tasks, []int{})

		if string(out) != expected {
			t.Errorf("expected %q - got %q", expected, string(out))
		}

		listcmd := exec.Command(path, "-list")

		out, err = listcmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected = stringer(tasks, []int{})

		if string(out) != expected {
			t.Errorf("expected %q - got %q", expected, string(out))
		}

		for i := 1; i <= len(tasks); i++ {
			cmd := exec.Command(path, []string{"-complete", strconv.Itoa(i)}...)

			_, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}
		}

		allcmd = exec.Command(path, "-all")

		out, err = allcmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected = stringer(tasks, []int{1, 2, 3, 4})

		if string(out) != expected {
			t.Errorf("expected %q - got %q", expected, string(out))
		}

		listcmd = exec.Command(path, "-list")

		out, err = listcmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected = ""

		if string(out) != expected {
			t.Errorf("expected %q - got %q", expected, string(out))
		}

		for i := 1; i <= len(tasks); i++ {
			uncmpcmd := exec.Command(path, []string{"-uncomplete", strconv.Itoa(i)}...)

			_, err = uncmpcmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}
		}

		cmd := exec.Command(path, "-list")

		out, err = cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected = stringer(tasks, []int{})

		if string(out) != expected {
			t.Errorf("expected %q - got %q", expected, string(out))
		}

		cmd = exec.Command(path, "-all")

		out, err = cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected = stringer(tasks, []int{})

		if string(out) != expected {
			t.Errorf("expected %q - got %q", expected, string(out))
		}
	})
}

func stringer(strs []string, dones []int) string {
	formatted := ""

	for i, str := range strs {
		prefix := "   "
		if isExist(i+1, dones) {
			prefix = "X  "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, i+1, str)
	}

	return formatted
}

func isExist(num int, numbers []int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}

	return false
}
