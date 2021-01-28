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
	"isready/pkg"

	"github.com/spf13/cobra"
)

// psqlCmd represents the psql command
var psqlCmd = &cobra.Command{
	Use:   "psql",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("psql called")

		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		host, _ := cmd.Flags().GetString("host")
		database, _ := cmd.Flags().GetString("database")
		port, _ := cmd.Flags().GetInt("port")

		conn := pkg.SQLConnection{
			Driver:   "postgres",
			Username: user,
			Password: password,
			Host:     host,
			Port:     port,
			Database: database,
			Timeout:  timeout,
			Retries:  int(retries),
		}

		err := pkg.OpenSQL(conn)
		if err != nil {
			fmt.Errorf("pg error: " + err.Error())
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	psqlCmd.Flags().String("user", "foo", "postgres usernamea")
	psqlCmd.Flags().String("password", "bar", "postgres password")
	psqlCmd.Flags().String("host", "localhost", "postgres host")
	psqlCmd.Flags().Int("port", 5432, "postgres database port")
	psqlCmd.Flags().String("database", "default", "name of the postgres database")

	rootCmd.AddCommand(psqlCmd)

}
