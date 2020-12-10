package tasks

import "time"

type Session struct {
	Date    time.Weekday
	Rdcount int
}
type Task struct {
	Subject  string
	Today    Session
	Sessions []Session
}

type Tasklist []Task
