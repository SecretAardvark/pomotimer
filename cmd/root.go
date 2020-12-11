package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//Interval is the time duration in minutes you wish to study for.
var Interval int

//Subject is the topic you wish to focus on.
var Subject string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomotimer",
	Short: "Pomotimer is a productivity tracker that uses the pomodoro method.",
	Long: `Pomotimer is a productivity tracker that uses the pomodoro method, a 
	time management methodology where you work for 25 minutes followed by a 5 minute
	break. 
	
	Usage examples: 
	
	Start command: Starts a pomodoro timer. Subject and timer interval can be set with 
	-s and -i flags.
		-pomotimer start -s thingToDo
		
	Add command: Adds a focus topic to the database. Requires the -s flag. 
		pomotimer add -s thingToDo

	Delete command: Removes a topic from the datagbase. Requires the -s flag. 
		pomotimer delete -s thingToDo
	
	Show command: Shows how much time you've spent on a given subject. Show will 
	display data for one topic if given the -s flag, defaulting to showing the whole 
	database. 
		pomotimer show -s thingToDo
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobratimer.yaml)")
	rootCmd.PersistentFlags().StringVarP(&Subject, "subject", "s", "", "The topic you wish to focus on.")
	rootCmd.PersistentFlags().IntVarP(&Interval, "interval", "i", 10, "The focus interval for the timer.")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobratimer" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobratimer")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
