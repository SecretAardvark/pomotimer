package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"pomotimer/tasks"
)

//OpenDB opens  the JSON database and writes it into the task list.
func OpenDB(taskList []tasks.Task) []tasks.Task {
	jsonfile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(errors.New("couldn't open the json file"))
	}
	defer jsonfile.Close()
	fmt.Println("Opened tasks.json...")
	byteValue, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		fmt.Println(errors.New("couldn't read the json file"))
	}
	json.Unmarshal(byteValue, &taskList)

	return taskList
}

//WriteDB writes the tasklist back into JSON.
func WriteDB(taskList []tasks.Task, task tasks.Task) {
	fmt.Println(taskList, task)
	newList := []tasks.Task{task}

	taskList = append(newList, taskList...)

	CloseDB(taskList)
}

//CloseDB closes the JSON file.
func CloseDB(taskList []tasks.Task) {
	file, _ := json.MarshalIndent(taskList, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
	os.Exit(1)
}
