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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	jira "github.com/andygrunwald/go-jira"
)

func SelectSprintInProject(jiraClient *jira.Client, boardID int) *jira.Sprint {

	sprints, _, _ := jiraClient.Board.GetAllSprintsWithOptions(boardID, &jira.GetAllSprintsOptions{State: "active,future"})

	for _, sprint := range sprints.Values {
		fmt.Printf(" Sprint ID : %d : %s\n", sprint.ID, sprint.Name)
	}

	r := bufio.NewReader(os.Stdin)

	fmt.Print("Please input desired sprint ID: ")
	sprintID, _ := r.ReadString('\n')
	sprintID = strings.TrimSuffix(sprintID, "\n")
	sprintID = strings.TrimSuffix(sprintID, "\r")

	// TODO Validate input

	for _, sprint := range sprints.Values {
		intSpID, _ := strconv.Atoi(sprintID)
		if sprint.ID == intSpID {
			return &sprint
		}
	}

	return nil

}

func SelectSprintListInProject(jiraClient *jira.Client, boardID int) []jira.Sprint {

	sprints, _, _ := jiraClient.Board.GetAllSprintsWithOptions(boardID, &jira.GetAllSprintsOptions{State: "active,future"})

	for _, sprint := range sprints.Values {
		fmt.Printf(" Sprint ID : %d : %s\n", sprint.ID, sprint.Name)
	}

	r := bufio.NewReader(os.Stdin)

	fmt.Print("Please input desired sprint IDs, comma separated: ")
	sprintIDs, _ := r.ReadString('\n')
	sprintIDs = strings.TrimSuffix(sprintIDs, "\n")
	sprintIDs = strings.TrimSuffix(sprintIDs, "\r")

	// TODO Validate input

	var selectedSprints []jira.Sprint
	for _, sprint := range sprints.Values {

		if strings.Contains(sprintIDs, strconv.Itoa(sprint.ID)) {
			selectedSprints = append(selectedSprints, sprint)
		}
	}

	return selectedSprints

}
