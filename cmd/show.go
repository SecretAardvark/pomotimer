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
//TODO: --date flag for show command to index the database by date?
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("show called")
		jsonFile, err := os.Open("test.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		bytes, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(bytes, &tasklist)
		if Subject == "" {
			fmt.Println(tasklist)
		}
		for _, s := range tasklist {
			if s.Subject == Subject {
				fmt.Println(s)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
