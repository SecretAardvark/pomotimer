package timer

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

//Start starts a pomodoro timer.
func Start(pause chan struct{}, interval int, elapsed time.Duration, startTime time.Time, subject string) {
	fmt.Printf("Starting the timer for subject: %s!\n", subject)
	timer := time.NewTimer(time.Duration(interval) * time.Second)
	select {
	case <-timer.C:
		beeep.Alert("Pomotimer", "Time's up!", "timer.png")
		fmt.Println("Study round is over.")
		elapsed = time.Now().Sub(startTime)
		StartBreak()
	case <-pause:
		timer.Stop()
		fmt.Println("Timer is paused.")
	}
}

//StartBreak starts the break timer.
func StartBreak() {
	fmt.Println("Starting the break timer!")
	breakTimer := time.NewTimer(5 * time.Second)
	<-breakTimer.C
	beeep.Alert("Pomotimer", "Time's up!", "timer.png")
	fmt.Println("Break is over!")
}
