/*
Copyright © 2020 Chad Tennent th3b0swick@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
//TODO: Implement changing the round interval to the interval flag.

//BUG: Using the start command with a subject that doesn't already exist overwrites the subject in test.json with just the single subject.
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*fmt.Println("start called")
		fmt.Println(Subject)
		fmt.Println(Interval)*/
		var tasklist tasks.Tasklist
		var timerStarted = false
		var elapsed time.Duration
		var startTime time.Time
		currentTask := tasks.Task{Subject: Subject}

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
				defer jsonFile.Close()
				for _, t := range tasklist {
					if currentTask.Subject == t.Subject {
						t.Sessions = append(t.Sessions, currentTask.Today)
						t.Today = currentTask.Today
						t.Today.Date = time.Now().Weekday()
					}
					byteValue, _ := ioutil.ReadAll(jsonFile)
					json.Unmarshal(byteValue, &tasklist)
					//tasklist = append(tasklist, currentTask)
					file, _ := json.MarshalIndent(tasklist, "", " ")
					_ = ioutil.WriteFile("test.json", file, 0644)
					os.Exit(1)
				}

				currentTask.Sessions = append(currentTask.Sessions, currentTask.Today)
				byteValue, _ := ioutil.ReadAll(jsonFile)
				json.Unmarshal(byteValue, &tasklist)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
