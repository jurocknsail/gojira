# GoJira
 [![Build Status](https://travis-ci.org/jurocknsail/gojira.svg?branch=master)](https://travis-ci.com/jurocknsail/gojira)

![gojira](https://github.com/jurocknsail/gojira/blob/master/logo/gojira-logo_150px.png)

## What is GoJira ?

GoJira is a Go tool designed to enhance and expedite the use of Jira, a popular management software.
This version of GoJira focuses on enabling Definition of Done (DoD) injection.

## How to install GoJira ?

#### Download binary

* Visit the [Releases page](https://github.com/gemalto/gokube/releases/latest) to access the latest GoJira release.
* Copy the executable file to a directory included in your system's PATH
* The executable file will be named as gojira-version-type+platform.arch.exe. For ease of use, rename it to "gojira.exe".

#### Verify the Executable

Verify the executable installation by opening your preferred command-line interface (CLI) and entering the command "gojira". You should see output similar to the following:

```shell
$ gojira.exe
Gojira is a modern CLI Tool for Jira Server, allowing DoD injection and plenty other stuff impossible to achieve using Jira plugins.

Usage:
  gojira [flags]
  gojira [command]

Available Commands:
  config      Manage gojira configuration.
  dod         Apply Definition of Done to a set of US.
  help        Help about any command
  login       Create wincred credential to authenticate to a dedicated Jira Server / Project via Basic Authentication.
  pi-wl       List all PI sprints workload
  sprints     List all sprints in configured project.
  stories     List all User Stories in the configured Sprint.
  test
  tokenlogin  Create wincred credential to authenticate to a dedicated Jira Server / Project via Bearer Token.
  version     Print the version of Gojira

Flags:
  -h, --help   help for gojira

Use "gojira [command] --help" for more information about a command.

```
If you see the above output, the installation is complete.
Otherwise, double-check that you placed the gojira.exe file in the correct directory and correctly added it to your system's PATH variable. 

#### Configure gojira

* Create Jira credentials: Set up your Jira credentials once to establish authentication.

```shell
$ gojira.exe login
Jira URL: https://jira.url.com
Jira Username: username
Jira Password: *********
Jira account was successfully logged on!
```
```shell
$ gojira.exe tokenlogin
Jira URL: https://jira.gemalto.com
Jira Personal Access Token:
Jira account was successfully logged on!
```

* Configure the Jira Project: Specify the Jira Project that you want to work with using GoJira.

```shell
$ gojira.exe config project "<MY_PROJECT_NAME>"
Several boards matching, please pick one by ID.
Board : {XXXXX https://jira.url.com/rest/agile/1.0/board/XXXXX MY_PROJECT_NAME scrum 0}
Board : {YYYYY https://jira.url.com/rest/agile/1.0/board/YYYYY OTHER_PROJECT_NAME scrum 0}
Selected Board ID: XXXXX
Board selected!
```

* Configure the Jira Project ID: You can also specify the Jira Project ID (if you have it) that you want to work with using GoJira.

```shell
$ gojira.exe config projectId <MY_PROJECT_ID>
Board selected!
Board ID: XXXXX
Board Name: MY_PROJECT_NAME
```

* Configure the desired sprint: Specify the GoJira configuration to use the desired sprint in the selected project. (Note: This configuration can be updated later by reusing the same command and selecting another sprint.)

```shell
$ gojira.exe config sprint
 Sprint ID : AAAAA : Sprint 1
 ...
 Sprint ID : CCCCC : Sprint 3
 ...
Please input desired sprint ID: CCCCC
```

#### Functionalities

##### List of User Stories

* It is possible to list all User Stories (US) present in the selected sprint, for example in order to select some we need to modify using another gojira command.

```shell
$ gojira.exe stories
 
Stories/Bugs in selected sprint :
 PROJ-DDDDD | Story  | A User Story
 PROJ-EEEEE | Story  | Another US   
 PROJ-FFFFF | Story  | Again a US   
 
As list :
PROJ-DDDDD,PROJ-EEEEE,PROJ-FFFFF
```

##### DoD Injection

* GoJira enables you to apply a desired Definition of Done to a list of User Stories.
* Usage : gojira dod [ dodname ] US-XXXXX,US-YYYYY,....

```shell
$ gojira.exe dod standardstory PROJ-DDDDD,PROJ-EEEEE
 Pushing standardstory DoD for US PROJ-DDDDD ...............
 Pushing standardstory DoD for US PROJ-EEEEE ...............
Number of US treated : 2
```

## Additional Links

* [**Contributing**](./CONTRIBUTING.md)
* [**Development Guide**](./docs/developer-guide.md)

