package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AliiAhmadi/td"
)

// File name
const todoFileName = ".todo.json"

func main() {
	changeUsage()

	// Define command line arguments
	all := flag.Bool("all", false, "List all tasks")
	task := flag.String("task", "", "Task to be included i the toDo list")
	list := flag.Bool("list", false, "List uncompleted tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	uncomplete := flag.Int("uncomplete", 0, "Uncomplete a completed task")

	flag.Parse()

	l := &td.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *all:
		// List all todos
		fmt.Print(l)

	case *list:
		// List just completed todos
		fmt.Print(l.Format())

	case *complete > 0:
		// Complete a task with this index
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Show updated task
		fmt.Fprintln(os.Stdout, (*l)[*complete-1].Task)

	case *task != "":
		// Creating new task
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *uncomplete > 0:
		// Uncomplete a task
		if err := l.Uncomplete(*uncomplete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Show updated task
		fmt.Fprintln(os.Stdout, (*l)[*uncomplete-1].Task)

	default:
		// Invalid option
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func changeUsage() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool, Developed for MHM\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2024\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}
}
