//TODO:  Add different focus subjects to JSON and track how much time is spent on each.
//TODO: CLI frontend for timer start/stop/tracking
//FIXED: elapsed always == 0?
//BUG: You can pause the timer before starting it.
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

	var rdCount int
	for scanner.Scan() {
		startTime := time.Unix(100, 100)
		elapsed := time.Since(startTime)
		timer := time.NewTimer(10 * time.Second)
		switch scanner.Text() {
		case "start":
			startTime = time.Now()
			timer.Reset(10 * time.Second)
			rdCount++
			go start(*timer, startTime, elapsed, rdCount)
		case "pause":
			if startTime == time.Unix(100, 100) {
				fmt.Println("You didn't  start the timer!")
			} else {
				pause(timer, startTime, elapsed, rdCount)

			}
		case "unpause":
			if startTime == time.Unix(100, 100) {
				fmt.Println("You didn't  start the timer!")
			} else {
				unpause(timer, startTime, elapsed, rdCount)
			}

		}
	}
}

func start(timer time.Timer, startTime time.Time, elapsed time.Duration, rdCount int) {

	fmt.Println("Starting timer")

	<-timer.C
	fmt.Println(time.Since(startTime))
	//Play system alert sound when timers alert.
	beeep.Notify("Timer", "Break Time!", "C:\\Users\\p\\Pictures\\fuck groupme\\lgbtq.jpg")
	timer.Reset(10 * time.Second)
	fmt.Println("Break Time!")
	fmt.Println(timer)
	breakTimer := time.NewTimer(5 * time.Second)
	<-breakTimer.C
	fmt.Println("Break is over")
	fmt.Println(rdCount)
	fmt.Println(time.Since(startTime))

}
func pause(timer *time.Timer, startTime time.Time, elapsed time.Duration, rdCount int) {
	fmt.Println("paused")
	timer.Stop()
	fmt.Printf("%v elapsed, %v remaining", time.Since(startTime),
		(10*time.Second - time.Since(startTime)))
}

func unpause(timer *time.Timer, startTime time.Time, elapsed time.Duration, rdCount int) {
	fmt.Println("unpaused")
	fmt.Printf("%v time remaining", (10*time.Second - time.Since(startTime)))
	timer.Reset(10*time.Second - time.Since(startTime))
}
