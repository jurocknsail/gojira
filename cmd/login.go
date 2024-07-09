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
	"strings"
	"syscall"

	"github.com/howeyc/gopass"
	"github.com/spf13/viper"

	jira "github.com/andygrunwald/go-jira"
	"github.com/danieljoos/wincred"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var gojiraCredentialsName = "gojira"
var gojiraAuthTransport = ""

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(loginTokenCmd)
	loginCmd.AddCommand(deleteCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Create wincred credential to authenticate to a dedicated Jira Server / Project via Basic Authentication.",
	Long:  `Create wincred credential to authenticate to a dedicated Jira Server / Project.via Basic Authentication.`,
	Run: func(cmd *cobra.Command, args []string) {

		gojiraAuthTransport = "basic"
		loginToJira()
	},
}

var loginTokenCmd = &cobra.Command{
	Use:   "tokenlogin",
	Short: "Create wincred credential to authenticate to a dedicated Jira Server / Project via Bearer Token.",
	Long:  `Create wincred credential to authenticate to a dedicated Jira Server / Project via Bearer Token.`,
	Run: func(cmd *cobra.Command, args []string) {

		gojiraAuthTransport = "bearer"
		loginToJira()
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete wincred credential to aunthenticate to a dedicated Jira Server / Project.",
	Long:  `Delete wincred credential to aunthenticate to a dedicated Jira Server / Project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deleting '%s' credential from wincred ... \n", gojiraCredentialsName)
		cred, err := wincred.GetGenericCredential(gojiraCredentialsName)
		if err != nil {
			fmt.Println(err)
			return
		}
		cred.Delete()
	},
}

func loginToJira() *jira.Client {

    var client *jira.Client
    var username string

	cred, err := wincred.GetGenericCredential(gojiraCredentialsName)
	if err == nil {
		// fmt.Printf("Logging as : %s\n", cred.Credential.UserName)
        if len(cred.Credential.UserName) == 0 {
            return createJiraClientToken(string(cred.Attributes[0].Value), string(cred.CredentialBlob))
        } else {
		    return createJiraClient(string(cred.Attributes[0].Value), cred.Credential.UserName, string(cred.CredentialBlob))
        }
	} else {
		//Create the credential
		r := bufio.NewReader(os.Stdin)

		fmt.Print("Jira URL: ")
		jiraURL, _ := r.ReadString('\n')
		jiraURL = strings.TrimSuffix(jiraURL, "\n")
		jiraURL = strings.TrimSuffix(jiraURL, "\r")

        if gojiraAuthTransport == "basic" {
            fmt.Print("Jira Username: ")
            username, _ = r.ReadString('\n')
            username = strings.TrimSuffix(username, "\n")
            username = strings.TrimSuffix(username, "\r")
            fmt.Print("Jira Password: ")
        } else{
		    fmt.Print("Jira Personal Access Token: ")
        }

		password, _ := gopass.GetPasswd()
		cred := wincred.NewGenericCredential(gojiraCredentialsName)
		cred.CredentialBlob = password

        if gojiraAuthTransport == "basic" {
            cred.UserName = username
		    client = createJiraClient(jiraURL, cred.Credential.UserName, string(cred.CredentialBlob))
        } else{
		    client = createJiraClientToken(jiraURL, string(cred.CredentialBlob))
        }

		// client.Board.GetAllBoards

		credAttributes := []wincred.CredentialAttribute{
			wincred.CredentialAttribute{
				"jiraUrl",
				[]byte(jiraURL),
			},
		}

		cred.Attributes = credAttributes
		cred.Persist = wincred.PersistEnterprise

		err := cred.Write()

		if err != nil {
			fmt.Println(err)
		}

        if gojiraAuthTransport == "basic" {
		    viper.Set("username", username)
        }
		viper.Set("jira_url", jiraURL)
		viper.WriteConfig()
		fmt.Printf("Jira account was successfully logged on!")

		return client
	}
}

func createJiraClient(jiraURL, username, password string) *jira.Client {

	r := bufio.NewReader(os.Stdin)

	if len(jiraURL) == 0 {
		fmt.Print("Jira URL: ")
		jiraURL, _ = r.ReadString('\n')
	}
	if len(username) == 0 {
		fmt.Print("Jira Username: ")
		username, _ = r.ReadString('\n')
	}

	if len(password) == 0 {
		fmt.Print("Jira Password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		password = string(bytePassword)
	}

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	jiraClient, err := jira.NewClient(tp.Client(), jiraURL)
	if err != nil {
		fmt.Printf("Error while creating Jira Client : %s\n", err)
		panic(err)
	}

	return jiraClient
}

func createJiraClientToken(jiraURL, password string) *jira.Client {

	r := bufio.NewReader(os.Stdin)

	if len(jiraURL) == 0 {
		fmt.Print("Jira URL: ")
		jiraURL, _ = r.ReadString('\n')
	}

	if len(password) == 0 {
		fmt.Print("Jira Personal Access Token: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		password = string(bytePassword)
	}

	tp := jira.BearerAuthTransport{
		Token: strings.TrimSpace(password),
	}

	jiraClient, err := jira.NewClient(tp.Client(), jiraURL)
	if err != nil {
		fmt.Printf("Error while creating Jira Client : %s\n", err)
		panic(err)
	}

	return jiraClient
}