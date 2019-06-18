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
	"fmt"
	"strconv"
	"strings"

	"github.com/andygrunwald/go-jira"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(piCmd)
	rootCmd.AddCommand(testCmd)
}

var piCmd = &cobra.Command{
	Use:   "pi-wl",
	Short: "List all PI sprints workload",
	Long:  `List all PI sprints workload`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {

		sprints := viper.GetString("pi_sprints")
		sprintList := strings.Split(sprints, ",")

		if sprintList != nil {
			client := loginToJira()

			//For each sprint
			for _, sprint := range sprintList {

				if sprint != "" {

					sprintSplit := strings.Split(sprint, "/")

					fmt.Printf("\nSprint : %s\n", sprintSplit[0])

					issuesFound, _ := listStoriesForSprint(client, sprintSplit[1])

					totalSP := 0
					for _, issue := range issuesFound {
						customFields, _, _ := client.Issue.GetCustomFields(issue.ID)
						storyPoints, _ := strconv.Atoi(customFields["customfield_10002"])
						// fmt.Printf(" %-11s | %-6s | %-150s | %-3d |\n", issue.Key, issue.Fields.Type.Name, issue.Fields.Summary, storyPoints)
						totalSP = totalSP + storyPoints
					}

					fmt.Printf("Total SP in sprint : %d\n", totalSP)
				}

			}

		} else {
			fmt.Printf("Please configure a PI sprints first using 'gojira config pi-sprints' command ! \n")
		}
	},
}

type config struct {
	BoardID    int           `yaml:"board_id"`
	BoardName  string        `yaml:"board_name"`
	JiraURL    string        `yaml:"jira_url"`
	PiSprint   string        `yaml:"pi_sprints"`
	SprintID   int           `yaml:"sprint_id"`
	SprintName string        `yaml:"sprint_name"`
	Username   string        `yaml:"username"`
	Test       []jira.Sprint `yaml:"test"`
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "",
	Long:  ``,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var conf config
		err := viper.Unmarshal(&conf)
		if err != nil {
			fmt.Printf("unable to decode into struct, %v", err)
		}
		fmt.Printf("%v", conf.Test[0].Name)
	},
}
