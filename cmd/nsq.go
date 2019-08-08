// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"ggcode/tpl"
	"github.com/spf13/cobra"
)

// nsqCmd represents the nsq command
var nsqCmd = &cobra.Command{
	Use:   "nsq",
	Short: "Create A Sample Nsq Server/Client",
	Long:  `Create Nsq code template. `,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		if !strings.HasSuffix(name, ".go") {
			name += ".go"
		}

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		source := fmt.Sprintf("%s/%s", dir, name)

		_isServer := cmd.Flag("server").Value.String()
		isServer, _ := strconv.ParseBool(_isServer)

		if isServer {
			// DO nothing
			return
		} else {
			if err := ioutil.WriteFile(source, []byte(tpl.NSQSERVER), 0644); err != nil {
				fmt.Println(err.Error())
			}
		}
	},
}

func init() {
	nsqCmd.PersistentFlags().Bool("server", false, "Create Server Code")
	nsqCmd.PersistentFlags().String("name", "nsq.go", "The NSQ Code File Name")
	rootCmd.AddCommand(nsqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nsqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nsqCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
