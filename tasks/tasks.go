package tasks

import "time"

//Session is your daily work session, today's date and how many study rounds you've done.
type Session struct {
	Date    time.Time
	Rdcount int
}

//Task is a subject you wish to focus on.
type Task struct {
	Subject  string    `json:"Subject"`
	Today    Session   `json:"Today"`
	Sessions []Session `json:"Sessions"`
}

//Tasklist is the master list of all your focus topics.
type Tasklist []Task
