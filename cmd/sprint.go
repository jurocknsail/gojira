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

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(sprintsCmd)
}

var sprintsCmd = &cobra.Command{
	Use:   "sprints",
	Short: "List all sprints in configured project.",
	Long:  `List all sprints in configured project.`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {

		boardID := viper.GetInt("board_id")
		if boardID != 0 {
			sprints, _, _ := loginToJira().Board.GetAllSprintsWithOptions(boardID, &jira.GetAllSprintsOptions{State: "active,future"})

			for _, sprint := range sprints.Values {
				fmt.Printf(" Sprint ID : %d : %s\n", sprint.ID, sprint.Name)
			}
		} else {
			fmt.Printf("Please configure a project first using 'gojira config project' command ! \n")
		}

	},
}

