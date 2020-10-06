//TODO: Find a more appropriate notification icon.
//TODO: Add the notification icon to the package.

//BUG: Pause function no longer stops the timer properly.

/*Architecture idea: Have Start/Pause/Unpause work in 3 seperate goroutines, and
send the timerStarted var across channels? */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"time"

	"github.com/SecretAardvark/pomotimer/tasks"
	"github.com/gen2brain/beeep"
	"github.com/thatisuday/commando"
)

var timerStarted = false

func main() {
	var taskList []tasks.Task
	//timerStarted == false
	//Take cli to start and stop the timer.
	commando.
		SetExecutableName("Pomotimer").
		SetVersion("1.0.0").
		SetDescription("Pomotimer is a productivity tracker using the pomodoro method.")

	commando.
		Register("start").
		SetDescription("Starts the pomodoro timer").
		AddFlag("task,t", "Set which task to focus on.", commando.String, "none").
		AddFlag("interval,i", "Set the focus duration in minutes", commando.Int, 25).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			if timerStarted == false {
				startTime := time.Now()
				var rdCount int
				timerStarted = true
				task := flags["task"].Value
				fmt.Println("Starting timer for task:", task)
				go start(startTime, rdCount)
			} else {
				fmt.Println("Timer is already running!")
			}
		})

	commando.
		Register("pause").
		SetDescription("Pauses the timer.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			if timerStarted == false {
				fmt.Println("You didn't start the timer!")
			} else {
				//pause()
			}
		})
	commando.
		Register("unpause").
		SetDescription("Unpauses the timer").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			if timerStarted == false {
				fmt.Println("You didn't start the timer!")
			} else {
				//unpause()
			}
		})
		// Add different focus subjects to JSON and track how much time is spent on each.
	commando.
		Register("add").
		AddArgument("task", "The task to add to the DB.", "").
		SetDescription("Adds a focus task to the DB.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			//Open the json file
			jsonFile, err := os.Open("test.json")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Opened Tasks.json")
			defer jsonFile.Close()

			//Read the JSON into a slice to work with
			byteValue, _ := ioutil.ReadAll(jsonFile)

			json.Unmarshal(byteValue, &taskList)

			//Add the task to the slice.
			fmt.Printf("Adding task:  '%s' to the DB.\n", args["task"].Value)
			newTask := tasks.Task{Name: args["task"].Value}
			taskList = append(taskList, newTask)
			fmt.Println(taskList)

			//Write the slice to back to JSON.
			file, _ := json.MarshalIndent(taskList, "", " ")
			_ = ioutil.WriteFile("test.json", file, 0644)

		})

	commando.
		Register("delete").
		AddArgument("task", "The task to be removed from the DB.", "none").
		SetDescription("Removes a task from the DB.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			//
			jsonFile, err := os.Open("test.json")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Opened tasks.json")
			defer jsonFile.Close()

			//Read the json into taskslist variable
			byteValue, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(byteValue, &taskList)

			//Remove the task.
			keys := make([]string, len(args))

			i := 0
			for k := range args {
				keys[i] = args[k].Value
				i++
			}
			fmt.Println(keys)

			for i, v := range keys {
				for _, task := range taskList {
					if v == task.Name {
						fmt.Printf("Deleting %v from the db\n", v)
						//taskList[i] = taskList[len(taskList)-1]
						//taskList[len(taskList)-1] = tasks.Task{}
						//taskList = taskList[:len(taskList)-1]
						copy(taskList[i:], taskList[i+1:])
						taskList[len(taskList)-1] = tasks.Task{}
						taskList = taskList[:len(taskList)-1]
					}
				}
				i++
			}

			//Save the tasklist back into JSON
			file, _ := json.MarshalIndent(taskList, "", " ")
			_ = ioutil.WriteFile("test.json", file, 0644)
			//Print call here just so we don't have to call "show" command for testing.
			fmt.Println(taskList)
		})
	commando.
		Register("show").
		AddArgument("task", "The task you wish to show stats for.", "none").
		SetDescription("Show stats for your tracked focus tasks.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			//Show function here.
			//Open the json file
			jsonFile, err := os.Open("test.json")
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()

			//Map the json to tasklist var
			bytes, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(bytes, &taskList)
			//Print it to the console
			fmt.Println(taskList)

		})

	commando.Parse(nil)

}

func start(startTime time.Time, rdCount int) (bool, *time.Timer) {
	fmt.Println("Starting timer")
	timer := time.NewTimer(10 * time.Second)
	timerStarted := true
	<-timer.C
	timerStarted = false
	fmt.Println(time.Since(startTime))
	//Play system alert sound when timers alert.
	beeep.Notify("Timer", "Break Time!", "C:\\Users\\p\\Pictures\\fuck groupme\\lgbtq.jpg")
	fmt.Println("Break Time!")
	breakTimer := time.NewTimer(5 * time.Second)
	<-breakTimer.C
	fmt.Println("Break is over")
	return timerStarted, timer
	//fmt.Println(rdCount)

}
func pause(timer *time.Timer, startTime time.Time, timerStarted bool) bool {
	fmt.Println("paused")
	timer.Stop()
	timerStarted = false
	fmt.Printf("%v elapsed, %v remaining", time.Since(startTime),
		(10*time.Second - time.Since(startTime)))
	return timerStarted
}

func unpause(timer *time.Timer, startTime time.Time, timerStarted bool) bool {
	fmt.Println("unpaused")
	fmt.Printf("%v time remaining", (10*time.Second - time.Since(startTime)))
	timer.Reset(10*time.Second - time.Since(startTime))
	timerStarted = true
	return timerStarted
}

/*
func remove(taskList []tasks.Task, task string) []tasks.Task {
	return append(slice :)
}*/
/*
func remove(taskList []tasks.Task, task int) []tasks.Task {
	ret := make([]tasks.Task, 0)
	ret = append(ret, taskList[:task]...)
	return append(ret, taskList[task+1:]...)
}*/
