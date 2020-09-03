package tasks

//Task represents different tasks you can focus on
//during pomodoro timer sessions.
type Task struct {
	name     string `json: Name`
	totalRds int    `json: TotalRds`
}

var Tasks []Task
