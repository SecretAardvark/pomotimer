package cmd

import (
	"fmt"
	"pomotimer/db"
	"pomotimer/tasks"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var tasklist tasks.Tasklist

		jsonDB := db.OpenDB(tasklist)

		if Subject != "" {
			for i, t := range jsonDB {
				if t.Subject == Subject {
					fmt.Printf("Deleting task %v from the db\n", t.Subject)
					a := jsonDB[:i]
					b := jsonDB[i+1:]
					jsonDB = append(a, b...)

					db.CloseDB(jsonDB)
					fmt.Println(tasklist) //Print tasklist here for testing.
					fmt.Println(jsonDB)
				}
				i++
			}

		}
		fmt.Println("Please specify a subject to delete.")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
