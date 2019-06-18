// Copyright 2019 Gemalto. All rights reserved.
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
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	helpers "github.com/gemalto/gojira/helpers"
)

var allowedDoDTypes = "feature | bug | sprint | pi | study | archi | vlr"

func init() {
	rootCmd.AddCommand(dodCmd)
	dodCmd.AddCommand(listCmd)
}

var dodCmd = &cobra.Command{
	Use:   "dod",
	Short: "Apply Definition of Done to a set of US.",
	Long:  `Apply Definition of Done to a set of US. The project and sprint used are the one configured using 'gojira config' commands.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Requires DoD type and a list of Stories")
		}
		if isValidDoDType(args[0]) {
			return nil
		}
		return fmt.Errorf("Invalid dod type: %s.\nDoD type should be in [ %s ]", args[0], allowedDoDTypes)
	},
	ValidArgs: []string{"feature", "bug", "sprint", "pi", "study", "archi", "vlr"},
	Example:   "gojira dod [ " + allowedDoDTypes + " ] US-XXXXX,US-YYYYY,....",
	Run: func(cmd *cobra.Command, args []string) {

		data, err := Asset("resources/subtasks.json")

		if err != nil {
			// Asset was not found.
		}
		sprintID := viper.GetString("sprint_id")
		username := viper.GetString("username")

		if sprintID != "" {

			jiraClient := loginToJira()
			issuesFound, _ := listStoriesForSprint(jiraClient, sprintID)

			usList := strings.Trim(strings.ToUpper(args[1]), " ")

			usInjected := 0
			for _, issue := range issuesFound {
				if strings.Contains(usList, issue.Key) {
					fmt.Printf(" Pushing %s DoD for US %s ", args[0], issue.Key)
					// use asset data
					var result map[string][]string
					json.Unmarshal([]byte(data), &result)
					for _, summary := range result[args[0]] {
						helpers.CreateSubTask(jiraClient, issue.Fields.Project.Key, username, issue.Key, issue.ID, summary)
						fmt.Printf(".")
					}
					usInjected = usInjected + 1
					fmt.Printf("\n")
				}
			}
			fmt.Printf("Number of US treated : %d\n", usInjected)

		} else {
			fmt.Printf("Please configure a sprint first using 'gojira config sprint' command ! \n")
		}

	},
}

func isValidDoDType(dodType string) bool {
	if strings.Contains(allowedDoDTypes, dodType) {
		return true
	}
	return false
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available DoD in config.",
	Long:  `List available DoD in config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		viper.SetConfigName("dod")
		viper.ReadInConfig()
		fmt.Printf("Configuration read from dod.yaml using viper : \n  %v \n\n", viper.AllKeys())
		fmt.Printf("Choosen dod : %v \n", args[0])

		fmt.Printf(" Subtasks : %v \n", viper.GetStringSlice(args[0]))

		//Restore config initial context
		viper.SetConfigName("gojira")
		viper.ReadInConfig()
	},
}
