package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	switch args[0] {
	case "start":
		Start()

		//Use flags to denote which task is being focused on.
	case "pause":
		fmt.Println("You haven't started a timer!")
	case "unpause":
		fmt.Println("You haven't started a timer!")
	case "add":
		//addTaskFunc()
		fmt.Println("Added a task(not really)")
	case "show":
		//Display tasks here
		//fmt.Println(Tasks )
		//Default to printing whole JSON, use flags for single tasks
	case "help":
		//Display small help doc.
		fmt.Println("Pomotimer Usage:\n start: start a pomodoro timer. \n \t-task use to denote the task you are focusing on. \nadd: adds a focus subject to the database. \nshow: Shows information about focus subjects. \n \t-subject shows a specific subject. \nhelp: Displays this help text. :)")
	}

}

func Start() {
	//	startTime := time.Now()
	timer := time.NewTimer(10 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Starting Timer")
	//	timerStarted := true
	switch scanner.Text() {
	case "pause":
		fmt.Println("Timer paused (not really)")
	case "unpause":
		//unpause func
	}
	<-timer.C
	fmt.Println("Timer finished.")

}
