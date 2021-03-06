/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var timeoutString string
var timeout time.Duration
var retries int32

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "is-ready",
	Short: "Check if service is ready",
	Long:  `isready is a powerful tool to check if a service is ready within a single command. This tool support a various collection of service to test with`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		timeout, err = time.ParseDuration(timeoutString)
		if err != nil {
			fmt.Errorf("could not parse timeout duration: " + err.Error())
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&timeoutString, "timeout", "30s", "timeout for connection")
	rootCmd.PersistentFlags().Int32Var(&retries, "retries", 3, "number of retries before abort")

}
