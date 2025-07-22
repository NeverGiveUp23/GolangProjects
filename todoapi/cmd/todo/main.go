package main

import (
	todo "felix/todoapi"
	"flag"
	"fmt"
	"os"
)

const todoFileName = ".todo.json"

func main() {
	//parsing command flags
	task := flag.String("-task", "", "Task to be included in the todo list")
	list := flag.Bool("-list", false, "List all tasks")
	complete := flag.Int("-complete", 0, "item to be completed")

	flag.Parsed()

	// define an item on the list
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//decide what to do based on the number of arguments provided
	switch {
	case *list:
		//List current to do items
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		//Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// add the task
		l.Add(*task)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
