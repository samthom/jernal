package main

import (
	"encoding/json"
	"os"
)

var filePath = "/Users/samthomas"
var fileName = "todo.json"

// Check if the todo.json file exists, if not create new file and inputs the root
// and returns the file interface
func CreateIfNotExist() (*os.File, error) {
    f, err := os.Open(filePath + fileName)
	if err != nil {
		fl, err := os.Create(filePath + "/" + fileName)
		if err != nil {
			return nil, err
		}
        // write root json to the file since its created newly
        err = writeRootToNewFile(fl)
		if err != nil {
			return nil, err
		}
		return fl, nil
	}
	return f, nil
}

func writeRootToNewFile(f *os.File) error {
    // create an empty root stuct instance
    tasks := []Task{}
    root := Root {
        Tasks: tasks,
    }

    return json.NewEncoder(f).Encode(root)
}
