package main

import (
	"encoding/json"
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

/* 
Example tasks list root json - 
{
    "tasks": [
        {
            "id": 102,
            "task": "wake up at 7am",
            "done": false
        },
        {
            "id": 103,
            "task": "sleep at 11pm",
            "done": false
        }
    ]
}
*/

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

type Root struct {
    Tasks []Task    `json:"tasks"`
}

func main() {
    argsLen := len(os.Args)
	if argsLen == 1 {
		printHelp()
		return
	}

	args := os.Args[1:]

	execute(args)
}

func execute(args []string) {
	f, err := CreateIfNotExist()
	if err != nil {
		log.Fatal(err)
	}
    defer f.Close()

	cmd := args[0]
	if cmd == list_cmd {
		listTaks(f)
	}
}

func listTaks(f *os.File) {
    // reading it all in one go, else we need to use buffered reader
    r, err := os.ReadFile(f.Name())
    if err != nil {
        fmt.Print(err)
        return
    }
    var root Root
    json.Unmarshal(r, &root)
    if len(root.Tasks) == 0 {
        fmt.Println("No tasks yet, use add to add tasks")
    }
    return
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
