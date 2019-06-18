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

	"github.com/spf13/viper"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(storiesCmd)
}

var storiesCmd = &cobra.Command{
	Use:   "stories",
	Short: "List all User Stories in the configured Sprint.",
	Long:  `List all User Stories in the configured Sprint.`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {

		sprintID := viper.GetString("sprint_id")
		if sprintID != "" {
			client := loginToJira()
			issuesFound, usAsList := listStoriesForSprint(client, sprintID)

			totalSP := 0
			fmt.Printf("\nStories/Bugs in selected sprint : \n")
			for _, issue := range issuesFound {
				customFields, _, _ := client.Issue.GetCustomFields(issue.ID)
				storyPoints, _ := strconv.Atoi(customFields["customfield_10002"])
				fmt.Printf(" %-11s | %-6s | %-150s | %-3d |\n", issue.Key, issue.Fields.Type.Name, issue.Fields.Summary, storyPoints)
				totalSP = totalSP + storyPoints
			}

			fmt.Printf("\nAs list :\n")
			fmt.Printf("%s\n", usAsList)
			fmt.Printf("\nTotal SP in sprint : %d\n", totalSP)

		} else {
			fmt.Printf("Please configure a sprint first using 'gojira config sprint' command ! \n")
		}
	},
}

func listStoriesForSprint(jiraClient *jira.Client, sprintID string) ([]jira.Issue, string) {
	issuesFound, _, _ := jiraClient.Issue.Search("Sprint = "+sprintID+" AND (issuetype = 'Story' OR issuetype = 'Bug')", &jira.SearchOptions{StartAt: 0, MaxResults: 50})

	usAsList := ""
	for _, issue := range issuesFound {
		usAsList = usAsList + "," + issue.Key
	}

	last := len(usAsList)
	if len(usAsList) > 0 {
		usAsList = usAsList[1:last]
	}

	return issuesFound, usAsList
}
