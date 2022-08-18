package main

import (
	"fmt"
	"log"
	"os"
)

const done = `[x]`
const notdone = `[ ]`

const (
	list_cmd = "list"
	add_cmd  = "add"
)

const help = `
	todo help
	add <task> - add task to the todo list.
	list - to list all the tasks.
	done <task id> - to mark a task as done.
	delete <task id> - to delete a task from the list

	`

type task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var taskList = []task{
	{
		ID:   1,
		Task: "wake up at 7.30",
		Done: false,
	},
	{
		ID:   2,
		Task: "sleep at 12",
		Done: false,
	},
	{
		ID:   3,
		Task: "eat healthy",
		Done: true,
	},
}

func main() {

	if len(os.Args) == 1 {
		printHelp()
		return
	}

	args := os.Args[1:]

	execute(args)
	err := CreateIfNotExist()
	if err != nil {
		log.Fatal(err)
	}
}

func execute(args []string) {
	cmd := args[0]
	if cmd == list_cmd {
		listTaks()
	}
}

func listTaks() {
	for _, v := range taskList {
		fmt.Printf("%v        %v. %v \n", isDone(v.Done), v.ID, v.Task)
	}
}

func isDone(status bool) string {
	if status {
		return done
	}
	return notdone
}

func printHelp() {
	fmt.Println(help)
}
