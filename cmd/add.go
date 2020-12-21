package cmd

import (
	"fmt"
	"os"
	"pomotimer/db"
	"pomotimer/tasks"

	"github.com/spf13/cobra"
)

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
		jsonDB := db.OpenDB(Tasklist)
		fmt.Println("Adding task to the db.")
		for _, task := range jsonDB {
			if Subject == task.Subject {
				fmt.Println("This task already exists.")
				os.Exit(1)
			}

		}
		task := tasks.Task{
			Subject: Subject,
		}

		db.WriteDB(jsonDB, task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
