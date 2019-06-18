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
	"io/ioutil"

	jira "github.com/andygrunwald/go-jira"
)

func CreateSubTask(jiraClient *jira.Client, jiraProject string, jiraUser string, parentIssueKey string, parentIssueID string, summary string) {

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: jiraUser,
			},
			Reporter: &jira.User{
				Name: jiraUser,
			},
			Type: jira.IssueType{
				Name: "Sub-task",
			},
			Project: jira.Project{
				Key: jiraProject,
			},
			Description: "Auto generated subtask from DoD",
			Summary:     summary,
			Parent: &jira.Parent{
				Key: parentIssueKey,
				ID:  parentIssueID,
			},
		},
	}
	_, resp, err := jiraClient.Issue.Create(&i)
	if resp.StatusCode != 201 {
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Error while creating subtask :  %s\n", b)
		panic(err)
	}
}
