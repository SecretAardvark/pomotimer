package timer

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

func Start(pause chan struct{}, interval int, elapsed time.Duration, startTime time.Time, subject string) {
	fmt.Printf("Starting the timer for subject: %s!", subject)
	timer := time.NewTimer(time.Duration(interval) * time.Second)
	select {
	case <-timer.C:
		beeep.Alert("Pomotimer", "Time's up!", "C:/Users/p/Pictures/GroupMe/GroupMe_2020724_133228.jpeg")
		fmt.Println("Study round is over.")
		elapsed = time.Now().Sub(startTime)
		StartBreak()
	case <-pause:
		timer.Stop()
		fmt.Println("Timer is paused.")
	}
}

func StartBreak() {
	fmt.Println("Starting the break timer!")
	breakTimer := time.NewTimer(5 * time.Second)
	<-breakTimer.C
	beeep.Alert("Pomotimer", "Time's up!", "C:/Users/p/Pictures/GroupMe/GroupMe_2020724_133228.jpeg")
	fmt.Println("Break is over!")
}
