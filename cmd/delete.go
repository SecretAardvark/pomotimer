//TODO: Make command usage uniform.
//		The delete command works as is but the usage is inconsistent with the other
//		commands. Delete takes an Arg for subject and Add command uses a flag.
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

		jsonfile, err := os.Open("test.Json")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Opened tasks.json")
		defer jsonfile.Close()

		byteValue, _ := ioutil.ReadAll(jsonfile)
		json.Unmarshal(byteValue, &tasklist)

		keys := make([]string, len(os.Args))

		i := 0
		for k := range os.Args {
			keys[i] = os.Args[k]
			i++
		}
		fmt.Println(keys)

		for _, v := range keys {
			for i, task := range tasklist {
				if v == task.Subject {
					fmt.Printf("Deleting %v from the db\n", v)
					a := tasklist[:i]
					b := tasklist[i+1:]
					tasklist = append(a, b...)

				}
			}
			i++
		}
		file, _ := json.MarshalIndent(tasklist, "", " ")
		_ = ioutil.WriteFile("test.json", file, 0644)
		fmt.Println(tasklist) //Print tasklist here for testing.

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
