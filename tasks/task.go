package tasks

import (
	"time"
)

type today time.Time

//Task represents different tasks you can focus on
//during pomodoro timer sessions.
type Task struct {
	name     string
	totalRds int
	Sessions map[time.Time]int
}

//Tasks is the master list of tasks you've worked on
var Tasks []Task
