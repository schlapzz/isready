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
	"encoding/json"
	"fmt"
	"isready/pkg"
	url2 "net/url"
	"os"

	"github.com/spf13/cobra"
)

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "curl",
	Short: "checks if a service is available with a http request",
	Long:  `checks if a service is available with a http request.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("curl called")

		host, _ := cmd.Flags().GetString("host")
		headers, _ := cmd.Flags().GetString("header")
		code, _ := cmd.Flags().GetInt32Slice("code")
		method, _ := cmd.Flags().GetString("method")
		insecure, _ := cmd.Flags().GetBool("insecure")

		url, err := url2.Parse(host)
		if err != nil {
			os.Stderr.WriteString("http error: " + err.Error())
			os.Exit(23)
		}

		header := make(map[string][]string)
		if headers != "" {
			err = json.Unmarshal([]byte(headers), header)
			if err != nil {
				os.Stderr.WriteString("could not parse http header parameter: " + err.Error())
				os.Exit(23)
			}
		}

		req := pkg.Http{
			Timeout:            timeout,
			Method:             method,
			Header:             header,
			URL:                url,
			StatusCodes:        code,
			SkipInsecureVerify: insecure,
		}

		err = req.Connect()
		if err != nil {
			os.Stderr.WriteString("http error: " + err.Error())
			os.Exit(23)
		}

	},
}

func init() {
	rootCmd.AddCommand(curlCmd)

	curlCmd.Flags().StringP("header", "h", "", "headers for http request")
	curlCmd.Flags().Int32Slice("code", nil, "accepted http response codes")
	curlCmd.Flags().String("host", "localhost", "url for http request")
	curlCmd.Flags().StringP("method", "X", "GET", "http request method")
	curlCmd.Flags().BoolP("insecure", "k", false, " By default, every SSL connection curl makes is verified to be secure. This option allows curl to proceed and operate even for server connections otherwise considered insecure.")

}
