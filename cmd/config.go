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
	"bytes"
	"errors"
	"fmt"
	"strconv"

	helpers "github.com/jurocknsail/gojira/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(useProjectCmd)
	configCmd.AddCommand(useProjectIdCmd)
	configCmd.AddCommand(useSprintCmd)
	configCmd.AddCommand(usePISprintsCmd)

}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage gojira configuration.",
	Long:  `Manage gojira configuration.`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		readerr := viper.ReadInConfig() // Find and read the config file
		if readerr != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %s", readerr))
		}
		fmt.Printf("Current GoJira configuration : \n")
		for _, key := range viper.AllKeys() {
			fmt.Printf(" %s: %s\n", key, viper.GetString(key))
		}
	},
}

var useProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create or update gojira configuration : project name used.",
	Long:  `Create or update gojira configuration : project name used.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires a project name, spaces must be escaped !")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		setProjectConfig(args[0])
	},
}

var useProjectIdCmd = &cobra.Command{
	Use:   "projectId",
	Short: "Create or update gojira configuration : project board id used.",
	Long:  `Create or update gojira configuration : project board id used.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires a project board id, spaces must be escaped !")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		setProjectConfigByID(args[0])
	},
}

var useSprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Create or update gojira configuration : sprint used.",
	Long:  `Create or update gojira configuration : sprint used.`,
	Run: func(cmd *cobra.Command, args []string) {
		setSprintConfig()
	},
}

var usePISprintsCmd = &cobra.Command{
	Use:   "pi-sprints",
	Short: "Create or update gojira configuration : sprints in PI.",
	Long:  `Create or update gojira configuration : sprints in PI.`,
	Run: func(cmd *cobra.Command, args []string) {
		setPISprintsConfig()
	},
}

func setProjectConfig(projectName string) {
	board := helpers.GetBoardByName(loginToJira(), projectName)
	if board == nil {
		fmt.Printf("No boards found, configuration has NOT been set ...\n")
	} else {
		viper.Set("board_name", board.Name)
		viper.Set("board_id", board.ID)
		viper.WriteConfig()
		fmt.Printf("Board selected! \n")
		fmt.Printf("Board ID: %s\n", strconv.Itoa(board.ID))
		fmt.Printf("Board Name: %s\n", board.Name)
	}
}

func setProjectConfigByID(projectID string) {

    projectBoardId, err := strconv.Atoi(projectID)
    if err != nil {
        fmt.Println("Error: incorrect value in projectID")
        return
    }
	board := helpers.GetBoardById(loginToJira(), projectBoardId)
	if board == nil {
		fmt.Printf("No boards found, configuration has NOT been set ...\n")
	} else {
		viper.Set("board_name", board.Name)
		viper.Set("board_id", board.ID)
		viper.WriteConfig()
		fmt.Printf("Board selected! \n")
		fmt.Printf("Board ID: %s\n", strconv.Itoa(board.ID))
		fmt.Printf("Board Name: %s\n", board.Name)
	}
}

func setSprintConfig() {
	//Get boardId from config
	boardID := viper.GetInt("board_id")
	if boardID != 0 {

		sprint := helpers.SelectSprintInProject(loginToJira(), boardID)
		if sprint != nil {
			viper.Set("sprint_name", sprint.Name)
			viper.Set("sprint_id", sprint.ID)
			viper.WriteConfig()
		} else {
			fmt.Printf("No such sprint found, configuration has NOT been set ...\n")
		}

	} else {
		fmt.Printf("Please configure a project first using 'gojira config project' or 'gojira config projectId' commands ! \n")
	}
}

func setPISprintsConfig() {
	//Get boardId from config
	boardID := viper.GetInt("board_id")
	if boardID != 0 {

		sprints := helpers.SelectSprintListInProject(loginToJira(), boardID)
		var selectedSprints bytes.Buffer
		for _, sprint := range sprints {
			selectedSprints.WriteString(sprint.Name)
			selectedSprints.WriteString("/")
			selectedSprints.WriteString(strconv.Itoa(sprint.ID))
			selectedSprints.WriteString(",")
		}

		if sprints != nil {
			viper.Set("pi_sprints", selectedSprints.String())
			viper.WriteConfig()
		} else {
			fmt.Printf("No such sprints found, configuration has NOT been set ...\n")
		}
	} else {
		fmt.Printf("Please configure a project first using 'gojira config project' or 'gojira config projectId' commands ! \n")
	}
}
