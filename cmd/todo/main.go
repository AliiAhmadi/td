package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AliiAhmadi/td"
)

// File name
const todoFileName = ".todo.json"

func main() {
	l := &td.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// with which flag what should happen ?
	switch {
	case len(os.Args) == 1:
		for _, todo := range *l {
			if !todo.Done {
				fmt.Fprintln(os.Stdout, todo.Task)
			}
		}

	default:
		if os.Args[1] == "-a" {
			for _, todo := range *l {
				fmt.Fprintln(os.Stdout, todo.Task)
			}
		} else {
			item := strings.Join(os.Args[1:], " ")
			l.Add(item)

			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
}
