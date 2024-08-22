package main

import (
	"flag"
	"fmt"
	"log"
)

func noArgsCheck(args []string) string {
	if len(args) == 0 {
		log.Fatal("Nothing to add")
	}
	value := args[0]
	if len(value) == 0 {
		log.Fatal("No value")
	}
	return value
}

func main() {
	addFlag := flag.Bool("add", false, "add todo flag")
	listFlag := flag.Bool("list", false, "list todo flag")
	doneFlag := flag.Bool("done", false, "mark as done todo flag")
	deleteFlag := flag.Bool("del", false, "mark as done todo flag")
	flag.Parse()
	switch {
	case *addFlag:
		{
			args := flag.Args()
			todoText := noArgsCheck(args)
			AddTodoHandler(todoText)
		}
	case *listFlag:
		ListAllTodo()
	case *doneFlag:
		args := flag.Args()
		id := noArgsCheck(args)
		MarkTodoDone(id)
	case *deleteFlag:
		args := flag.Args()
		id := noArgsCheck(args)
		DeleteTodoHandler(id)
	default:
		{
			fmt.Println("no block")
		}
	}
}
