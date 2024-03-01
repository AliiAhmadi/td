package main

import (
	"fmt"
	"os"

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
}
