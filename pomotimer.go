//TODO:  Add different focus subjects to JSON and track how much time is spent on each.
//TODO: CLI frontend for timer start/stop/tracking
//TODO: Play system alert sound when timers alert.
//BUG: Timer goroutine only runs once, never notifies when it's done after the first rd.
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	//Take cli to start and stop the timer.
	scanner := bufio.NewScanner(os.Stdin)
	startTime := time.Now()
	timer := time.NewTimer(10 * time.Second)
	elapsed := time.Since(startTime)

	var rdCount int
	for scanner.Scan() {
		switch scanner.Text() {
		case "start":
			timer.Reset(10 * time.Second)
			rdCount++
			go start(*timer, startTime, elapsed, rdCount)
		case "pause":
			fmt.Println("paused")
			timer.Stop()
			fmt.Println(timer)
			fmt.Println(elapsed)
		case "unpause":
			fmt.Println("unpaused")
			timer.Reset(10*time.Second - elapsed)
		}
	}
}

func start(timer time.Timer, startTime time.Time, elapsed time.Duration, rdCount int) {

	fmt.Println("Starting timer")

	<-timer.C
	fmt.Println(elapsed)
	beeep.Notify("Timer", "Break Time!", "C:\\Users\\p\\Pictures\\fuck groupme\\lgbtq.jpg")
	timer.Reset(10 * time.Second)
	fmt.Println("Break Time!")
	fmt.Println(timer)
	breakTimer := time.NewTimer(5 * time.Second)
	<-breakTimer.C
	fmt.Println("Break is over")
	fmt.Println(rdCount)

}
