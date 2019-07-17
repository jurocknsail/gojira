# GoJira
 [![Build Status](https://travis-ci.org/jurocknsail/gojira.svg?branch=master)](https://travis-ci.com/gemalto/gokube)

![gojira](https://github.com/jurocknsail/gojira/blob/master/logo/gojira-logo_150px.png)

## What is GoJira ?

GoJira is a Go tool meant to help & speed up our Jira usage.

The first version is dedicated to DoD injection.


## How to install GoJira ?

#### Download binary

* The latest release for gojira can be download on the [Releases page]<!-- (https://github.com/gemalto/gokube/releases/latest). -->
* Copy executable file to any directory which is in your PATH
* The gojira executable will be named as gojira-version-type+platform.arch.exe. Rename the executable to gojira.exe for ease of use.

#### Verify the Executable

In your preferred CLI, at the prompt, type gojira and press the Enter key. You should see output that starts with:

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
  login       Create wincred credential to aunthenticate to a dedicated Jira Server / Project.
  stories     List all User Stories in the configured Sprint.
  version     Print the version of Gojira

Flags:
  -h, --help   help for gojira

Use "gojira [command] --help" for more information about a command.

```
If you do, then the installation is complete.

If you don’t, double-check the path that you placed the gokube.exe file in and that you typed that path correctly when you added it to your PATH variable.

#### Configure gojira

* Create your jira credentials once

```shell
$ gojira.exe login
Jira URL: https://jira.url.com
Jira Username: username
Jira Password: *********
```

* Configure the Jira Project to be used

```shell
$ gojira.exe config project <MY_PROJECT_NAME>
Several boards matching, please pick one by ID.
Board : {XXXXX https://jira.url.com/rest/agile/1.0/board/XXXXX MY_PROJECT_NAME scrum 0}
Board : {YYYYY https://jira.url.com/rest/agile/1.0/board/YYYYY OTHER_PROJECT_NAME scrum 0}
Selected Board ID: XXXXX
```

* Configure gojira to use the desired sprint in selected project (can be changed later by reusing the same command and chosing another sprint)

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

* It is possible to list all US present in the selected sprint, for example in order to select some we need to modify using another gojira command.

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

* Apply the desired DoD to a list of US
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

