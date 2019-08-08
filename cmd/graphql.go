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
	"strings"

	"ggcode/tpl"
	"github.com/spf13/cobra"
)

// graphqlCmd represents the graphql command
var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "GraphQL Code",
	Long:  `Generate GraphQL Code. Default Port 80`,
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

		if err := ioutil.WriteFile(source, []byte(tpl.GRQPHQL), 0644); err != nil {
			fmt.Println(err.Error())
		}

		return

	},
}

func init() {
	rootCmd.AddCommand(graphqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	graphqlCmd.PersistentFlags().String("name", "graph.go", "The GraphQL Code File Name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
