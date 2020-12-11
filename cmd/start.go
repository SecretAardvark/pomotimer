//TODO: Fix the bufio scanner loop so you can 'start' a second round.
//BUG: On quit, Start command adds a duplicate entry to the db instead of overwriting an already existing one.
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"pomotimer/tasks"
	"pomotimer/timer"
	"time"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a pomodoro timer",
	Long: `The start command starts a pomodoro timer for the subject given by 
	the --subject flag, for a time interval set by the --interval flag. Defaults
	to no subject and 25 minute interval. 
	
	Examples: 
		Start a timer for subject 'focusTopic' for the default time interval: 
		pomotimer start -s focusTopic 
		pomotimer start --subject "focusTopic" 
		
		Start a timer for subject focusTopic with a custom timer interval: 

		pomotimer start -s focusTopic -i 15
		pomotimer start -s focusTopic --interval 15


	`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasklist tasks.Tasklist
		var timerStarted = false
		var elapsed time.Duration
		var startTime time.Time
		currentTask := tasks.Task{
			Subject: Subject,
			Today: tasks.Session{
				Date: time.Now().Local().UTC(),
			},
		}

		scanner := bufio.NewScanner(os.Stdin)
		pause := make(chan struct{}, 0)

		startTime = time.Now()
		timerStarted = true
		elapsed = 0 * time.Second
		currentTask.Today.Rdcount++
		go timer.Start(pause, Interval, elapsed, startTime, Subject)
		for scanner.Scan() {
			switch scanner.Text() {
			case "pause":
				if timerStarted == true && elapsed == 0*time.Second {
					pause <- struct{}{}
					timerStarted = false
					elapsed = time.Now().Sub(startTime)
					fmt.Printf("%v time elapsed", elapsed)
				}
			case "unpause":
				if timerStarted != true && elapsed > 0*time.Second {
					newTime := time.Duration(Interval) - elapsed
					go timer.Start(pause, int(newTime), elapsed, startTime, Subject)
				}
			case "quit":
				fmt.Println("See you next time!")
				jsonFile, err := os.Open("test.json")
				if err != nil {
					fmt.Println("Couldn't open the JSON file.")
				}
				byteValue, _ := ioutil.ReadAll(jsonFile)
				json.Unmarshal(byteValue, &tasklist)
				defer jsonFile.Close()

				for _, t := range tasklist {
					if currentTask.Subject == t.Subject {
						t.Sessions = append(t.Sessions, currentTask.Today)
						t.Today = currentTask.Today
						t.Today.Date = time.Now().Local().UTC()
						file, _ := json.MarshalIndent(tasklist, "", " ")
						_ = ioutil.WriteFile("test.json", file, 0644)
						os.Exit(1)
					}
				}

				currentTask.Sessions = append(currentTask.Sessions, currentTask.Today)
				tasklist = append(tasklist, currentTask)
				file, _ := json.MarshalIndent(tasklist, "", " ")
				_ = ioutil.WriteFile("test.json", file, 0644)
				os.Exit(1)

			}
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
