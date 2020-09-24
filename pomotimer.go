//TODO:  Add different focus subjects to JSON and track how much time is spent on each.
//TODO: CLI frontend for timer start/stop/tracking
//TODO: Find a more appropriate notification icon.
//TODO: Add the notification icon to the package.

//BUG: Pause function no longer stops the timer properly.

/*Architecture idea: Have Start/Pause/Unpause work in 3 seperate goroutines, and
send the timerStarted var across channels? */
package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/thatisuday/commando"
)

var timerStarted = false

func main() {
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
			fmt.Printf("Printing options of the base command.. \n\n")

			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
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
	commando.
		Register("add").
		AddArgument("task", "The task to add to the DB.", "").
		SetDescription("Adds a focus task to the DB.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Adding task:  '%s' to the DB.", args["task"].Value)
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
