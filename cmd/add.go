package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"pomotimer/tasks"

	"github.com/spf13/cobra"
)

var tasklist tasks.Tasklist

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a focus subject to the database.",
	Long: `Pomotimer is a productivity tracker that uses the pomodoro method. The add 
	command adds a focus subject you wish to track to the database. 
	
	Usage: 
	 	pomotimer add -s something
	  	pomotimer add --subject "anotherThing"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		jsonFile, err := os.Open("test.json")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Opened tasks.json")
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &tasklist)
		fmt.Println("Adding task to the db.")
		for _, task := range tasklist {
			if Subject == task.Subject {
				fmt.Println("This task already exists.")
				os.Exit(1)
			}

		}
		task := tasks.Task{
			Subject: Subject,
		}
		tasklist = append(tasklist, task)

		file, _ := json.MarshalIndent(tasklist, "", " ")
		_ = ioutil.WriteFile("test.json", file, 0644)
		fmt.Println(tasklist) //Print here for testing.
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
