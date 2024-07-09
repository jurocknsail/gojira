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

	"github.com/andygrunwald/go-jira"
)

func GetBoardByName(client *jira.Client, name string) *jira.Board {

    if client == nil {
       fmt.Println("Invalid Jira client! Please check Jira Credentials.")
       return nil
    }

	boards, _, err := client.Board.GetAllBoards(&jira.BoardListOptions{Name: name})
	if err != nil {
		fmt.Printf("Failed to get boards: %v\n", err)
		return nil
	}

	r := bufio.NewReader(os.Stdin)

	if len(boards.Values) > 1 {
		fmt.Printf("Several boards matching, please pick one by ID.\n")
		for _, board := range boards.Values {
			fmt.Printf("Board : %v\n", board)
		}
		fmt.Printf("Selected Board ID: ")

		boardID, _ := r.ReadString('\n')
		boardID = strings.TrimSuffix(boardID, "\n")
		boardID = strings.TrimSuffix(boardID, "\r")

		intBoardID, _ := strconv.Atoi(boardID)
		board, _, _ := client.Board.GetBoard(intBoardID)
		if board != nil {
			return board
		} else {
			fmt.Printf("No matching board, please try again.\n")
			return nil
		}

    } else if len(boards.Values) == 0 {
        fmt.Printf("No boards found. Please check Project Name parameter.\n")
        return nil

    } else {
        return &boards.Values[0]
    }

}

func GetBoardById(client *jira.Client, boardID int) *jira.Board {

    if client == nil {
       fmt.Println("Invalid Jira client! Please check Jira Credentials.")
       return nil
    }

	board, _, _ := client.Board.GetBoard(boardID)
	if board != nil {
        return board
	} else {
	    fmt.Printf("No matching board. Please check Project Board ID parameter.\n")
		return nil
	}

}