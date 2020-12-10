/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	 "Pomotimer add -s something
	  Pomotimer add --subject "anotherThing"`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
