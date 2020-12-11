package tasks

import "time"

//Session is your daily work session, today's date and how many study rounds you've done.
type Session struct {
	Date    time.Time
	Rdcount int
}

//Task is a subject you wish to focus on.
type Task struct {
	Subject  string
	Today    Session
	Sessions []Session
}

//Tasklist is the master list of all your focus topics.
type Tasklist []Task
